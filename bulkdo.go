package bulkdo

import (
	"bytes"
	"encoding/csv"
	"html/template"
	"io"
	"os"
	"os/exec"
)

func readItems(in io.Reader) ([]map[string]string, error) {
	reader := csv.NewReader(in)
	rows := make([]map[string]string, 0)
	var header []string
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		if header == nil {
			// 读 title 部分
			header = record
			continue
		}
		// 读 data 部分
		d := map[string]string{}
		for i := range header {
			d[header[i]] = record[i]
		}
		rows = append(rows, d)
	}
	return rows, nil
}

/*
echo {{.v.t1} {{.v.t2}}}
*/
func parseCommands(tplReader io.Reader, items []map[string]string) ([]string, error) {
	data, readErr := io.ReadAll(tplReader)
	if readErr != nil {
		return nil, readErr
	}
	t, parseErr := template.New("Commands").Parse(string(data))
	if parseErr != nil {
		return nil, parseErr
	}
	commands := []string{}
	for _, item := range items {
		p := make(map[string]map[string]string, 0)
		p["v"] = item

		var b bytes.Buffer
		exeErr := t.Execute(&b, p)
		if exeErr != nil {
			return nil, exeErr
		}
		commands = append(commands, b.String())

	}
	return commands, nil
}

func execCommands(cmds []string) ([]string, error) {
	outs := make([]string, 0)
	for _, cmd := range cmds {
		file, err := os.CreateTemp("", "bulkdo.*.bat") // 创建临时文件
		if err != nil {
			return nil, err
		}
		defer os.Remove(file.Name())

		file.WriteString(cmd)
		command := exec.Command("cmd.exe", "/c", file.Name())
		out, exeErr := command.CombinedOutput()
		if exeErr != nil {
			return nil, exeErr
		}
		outs = append(outs, string(out))
	}
	return outs, nil
}

// BulkDo is entry of bulkdo
func BulkDo(cmdTemplate io.Reader, params io.Reader) ([]string, error) {
	items, itemsErr := readItems(params)
	if itemsErr != nil {
		return nil, itemsErr
	}
	cmds, parseErr := parseCommands(cmdTemplate, items)
	if parseErr != nil {
		return nil, parseErr
	}
	outs, exeErr := execCommands(cmds)
	if exeErr != nil {
		return nil, exeErr
	}
	return outs, nil
}
