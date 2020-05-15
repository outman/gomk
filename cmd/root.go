package cmd

import (
	"fmt"
	"os"

	"github.com/gomarkdown/markdown"
	"github.com/spf13/cobra"
	"html/template"
	"io/ioutil"
)

var templateFile string
var input string
var output string

type MarkdownHtmlTemplate struct {
	Content template.HTML
}

var rootCmd = &cobra.Command{
	Use:   "gomk",
	Short: "Markdown to html with template",
	Long:  `A command tool for markdown to html with template based on gomarkdown/markdown`,
	Run: func(cmd *cobra.Command, args []string) {

		file, err := os.Stat(input)
		if err != nil {
			fmt.Println(err)
			return
		}

		switch mode := file.Mode(); {
		case mode.IsDir():
			fmt.Println("input markdown invalid.")
			return
		case mode.IsRegular():
			content, err := ioutil.ReadFile(input)
			if err != nil {
				fmt.Println(err)
				return
			}
			html := markdown.ToHTML(content, nil, nil)

			if templateFile != "" {
				_, err := os.Stat(templateFile)
				if err != nil {
					fmt.Println(err)
					return
				}
				tpl := template.Must(template.ParseFiles(templateFile))
				fileWriter, err := os.OpenFile(output, os.O_RDWR|os.O_CREATE, 0755)
				if err != nil {
					fmt.Println(err)
					return
				}
				defer fileWriter.Close()

				data := MarkdownHtmlTemplate{
					Content: template.HTML(string(html)),
				}
				tpl.Execute(fileWriter, data)
			} else {
				ioutil.WriteFile(output, html, 0644)
			}
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&templateFile, "template", "t", "", "your template path.")
	rootCmd.Flags().StringVarP(&output, "output", "o", "default.html", "your output file path name.")
	rootCmd.Flags().StringVarP(&input, "input", "i", "", "input markdown file path.")
	rootCmd.MarkFlagRequired("input")
}
