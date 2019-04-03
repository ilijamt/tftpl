package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"io/ioutil"
	"os"
	"text/template"
)

var buildCmd = &cobra.Command{
	Use:   "build [template-file]",
	Short: "Builds the file based on the specified template",
	Long:  `Builds the file based on the specified template`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var dataFile string
		var tplFile string
		var err error

		dataFile, _ = cmd.Flags().GetString("template-data")
		tplFile = args[0]

		var tplData string
		var rc io.ReadCloser
		if rc, err = openFile(tplFile); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer rc.Close()

		var buffer []byte
		if buffer, err = ioutil.ReadAll(rc); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		tplData = string(buffer)

		var tpl *template.Template
		if tpl, err = template.New("").Funcs(funcMap).Parse(tplData); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		var data interface{}
		if data, err = readDataFile(dataFile); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if err = tpl.Execute(os.Stdout, data); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

	},
}

func init() {
	rootCmd.AddCommand(buildCmd)

	flags := buildCmd.Flags()

	flags.StringP("template-data", "f", "", "")
	_ = cobra.MarkFlagRequired(flags, "template-data")
}
