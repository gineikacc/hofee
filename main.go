package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

var engine_desc = map[string]string{
	"fr":   "DeepL French",
	"fren": "DeepL French (reverse)",
	"ufr":  "YouGlish French",
	"wfr":  "Wiktionary French",
	// "wen":  "Wiktionary",
	// "pt":   "DeepL Portuguese",
	"wpt":  "Wiktionary Portuguese",
	"sp":   "DeepL Spanish",
	"spen": "DeepL Spanish (reverse)",
	"wsp":  "Wiktionary Spanish",
	"yt":   "Youtube",
}
var engine_uri = map[string]string{
	"wfr":  "https://en.wiktionary.org/wiki/{}#French",
	"fr":   "https://www.deepl.com/en/translator#en/fr/{}",
	"ufr":  "https://youglish.com/pronounce/{}/french/fr",
	"fren": "https://www.deepl.com/en/translator#fr/en/{}",
	// "wen":  "https://en.wiktionary.org/wiki/{}#English",
	"sp":   "https://www.deepl.com/en/translator#en/es/{}",
	"spen": "https://www.deepl.com/en/translator#es/en/{}",
	"wsp":  "https://en.wiktionary.org/wiki/{}#Spanish",
	// "pt":   "https://www.deepl.com/en/translator#en/pt/{}",
	"yt": "https://www.youtube.com/results?search_query={}",
}

func main() {
	//Ask for engine
	engine_used_str := run_rofi(true)
	engine_key := strings.Split(engine_used_str, "\t")[0]
	fmt.Println(engine_key)

	//Ask for prompt
	prompt := run_rofi(false)
	prompt = strings.ToLower(prompt)
	fmt.Println(prompt)

	//Launch engine w/ prompt
	run_engine(engine_uri[engine_key], prompt)

}

func run_rofi(ask_for_engine bool) string {

	cmd := exec.Command("rofi")
	cmd.Args = append(cmd.Args, "-dmenu")
	cmd.Args = append(cmd.Args, "-i")

	stdin, err := cmd.StdinPipe()
	if err != nil {
		panic(err)
	}

	var stdout bytes.Buffer
	cmd.Stdout = &stdout

	if err := cmd.Start(); err != nil {
		panic(err)
	}

	if ask_for_engine {
		var input []string
		for k, v := range engine_desc {
			line := fmt.Sprintf("%s\t%s", k, v)
			input = append(input, line)
		}
		_, err = stdin.Write([]byte(strings.Join(input, "\n")))
		if err != nil {
			panic(err)
		}
	}
	stdin.Close()

	if err := cmd.Wait(); err != nil {
		panic(err)
	}

	return stdout.String()
}

func run_engine(url, prompt string) {

	cmd := exec.Command("xargs")
	cmd.Args = append(cmd.Args, "-I{}")
	cmd.Args = append(cmd.Args, "xdg-open")
	cmd.Args = append(cmd.Args, url)

	stdin, err := cmd.StdinPipe()
	if err != nil {
		panic(err)
	}

	if err := cmd.Start(); err != nil {
		panic(err)
	}

	_, err = stdin.Write([]byte(prompt))
	if err != nil {
		panic(err)
	}
	stdin.Close()

	println("DEBUG")
	println(url)
	println(prompt)
	println("DEBUG OVER")
	if err := cmd.Wait(); err != nil {
		panic(err)
	}

}

//Ask which query engine to use
//Ask what prompt to fill
