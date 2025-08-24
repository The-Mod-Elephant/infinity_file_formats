package bg

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"slices"
	"strings"
)

type Component struct {
	TpFile        string `json:"tp_file"`
	Name          string `json:"name"`
	Lang          string `json:"lang"`
	Component     string `json:"component"`
	ComponentName string `json:"component_name"`
	SubComponent  string `json:"sub_component"`
	Version       string `json:"version"`
}

type Log struct {
	Comments   []string    `json:"comments"`
	Components []Component `json:"components"`
}

func OpenLog(r io.ReadSeeker) (*Log, error) {
	log := &Log{}

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		if !strings.HasPrefix(line, "~") {
			log.Comments = append(log.Comments, line)
		}
		component := Component{}
		parts := strings.Split(line, "~")
		// First part
		if len(parts) != 3 {
			continue
		}
		tpFileAndName := strings.Split(parts[1], string(os.PathSeparator))
		if len(tpFileAndName) != 2 {
			continue
		}
		component.Name = tpFileAndName[0]
		component.TpFile = tpFileAndName[1]
		langComponent := strings.Split(parts[2], "#")
		if len(langComponent) != 3 {
			continue
		}
		component.Lang = strings.TrimSpace(langComponent[1])
		componentSubComponent := strings.Split(langComponent[2], "// ")
		if len(componentSubComponent) != 2 {
			continue
		}
		component.Component = strings.TrimSpace(componentSubComponent[0])
		tail := strings.Split(componentSubComponent[1], " -> ")
		if len(tail) == 2 {
			component.ComponentName = tail[0]
			tail = strings.Split(tail[1], ": ")
			if len(tail) == 2 {
				component.SubComponent = tail[0]
				component.Version = tail[1]
			} else {
				component.SubComponent = strings.Join(tail, "")
			}
		} else {
			tail = strings.Split(tail[0], ": ")
			if len(tail) == 2 {
				component.ComponentName = tail[0]
				component.Version = tail[1]
			} else {
				component.ComponentName = strings.Join(tail, "")
			}
		}

		log.Components = append(log.Components, component)
	}
	return log, nil
}

func (l *Log) Write(w io.Writer) error {
	writer := bufio.NewWriter(w)
	for _, c := range l.Comments {
		writer.WriteString(c)
	}
	for _, c := range l.Components {
		out := fmt.Sprintf("~%s%s%s~ #%s #%s // %s", c.Name, string(os.PathSeparator), c.TpFile, c.Lang, c.Component, c.ComponentName)
		if c.SubComponent != "" {
			out = fmt.Sprintf("%s -> %s", out, c.SubComponent)
		}
		if c.Version != "" {
			out = fmt.Sprintf("%s: %s", out, c.Version)
		}
		writer.WriteString(out)
	}

	return nil
}

func (l *Log) WriteJson(w io.Writer) error {
	bytes, err := json.MarshalIndent(l, "", "\t")
	if err != nil {
		return err
	}

	_, err = w.Write(bytes)
	return err
}

func (l *Log) Equal(other *Log) bool {
	if !slices.Equal(l.Comments, other.Comments) {
		return false
	}
	if !slices.Equal(l.Components, other.Components) {
		return false
	}
	return true
}
