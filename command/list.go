package command

import (
	"archive/zip"
	"encoding/csv"
	"log"
	"os"
	"strings"

	"go-javap/parser"

	"github.com/urfave/cli"
)

func listCommand() cli.Command {
	return cli.Command{
		Name: "list",
		Action: func(c *cli.Context) error {
			w := csv.NewWriter(os.Stdout)
			w.Write([]string{"file", "class_type", "name", "super_name", "interfaces"})
			for _, file := range c.Args() {
				r, err := zip.OpenReader(file)
				if err != nil {
					return err
				}
				defer r.Close()
				for _, entry := range r.File {
					if !strings.HasSuffix(entry.Name, ".class") {
						continue
					}
					if entry.FileInfo().IsDir() {
						continue
					}
					entryReader, err := entry.Open()
					if err != nil {
						log.Printf("cannot open class file.  classfile=%s, jar=%s err=%v", entry.Name, file, err)
						return err
					}

					c, err := parser.ReadClass(entryReader)
					if err != nil {
						log.Printf("failed to parse %s: %v", entry.Name, err)
						continue
					}
					record := make([]string, 0)
					record = append(record, file)
					{
						var t string
						switch {
						case c.IsInterface():
							t = "interface"
						case c.IsAnnotation():
							t = "annotation"
						case c.IsEnum():
							t = "enum"
						case c.IsAbstract():
							t = "abstract"
						default:
							t = "class"
						}
						record = append(record, t)
					}
					record = append(record, c.Name())
					record = append(record, c.SuperClassName())
					record = append(record, strings.Join(c.Interfaces(), ", "))

					w.Write(record)
					w.Flush()
				}

			}
			return nil
		},
	}
}
