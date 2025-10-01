package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

var engine_desc = map[string]string{
	"fr":   "DeepL FR",
	"fren": "DeepL FR (reverse)",
	"ufr":  "YouGlish FR",
	"wfr":  "Wiktionary FR",
	"wen":  "Wiktionary",
	"uen":  "Wiktionary",
	// "pt":   "DeepL Portuguese",
	"wpt":  "Wiktionary PT",
	"sp":   "DeepL SP",
	"spen": "DeepL SP (reverse)",
	"wsp":  "Wiktionary SP",
	"ru":   "DeepL RU",
	"ruen": "DeepL RU (reverse)",
	"uru":  "YouGlish RU",
	"wru":  "Wiktionary RU",
	"nl":   "DeepL NL",
	"nlen": "DeepL NL (reverse)",
	"unl":  "YouGlish NL",
	"wnl":  "Wiktionary NL",
	"ge":   "DeepL GE",
	"geen": "DeepL GE (reverse)",
	"uge":  "YouGlish GE",
	"wge":  "Wiktionary GE",
	"lt":   "DeepL LT",
	"lten": "DeepL LT (reverse)",
	"wlt":  "DeepL LT",
	"yt":   "Youtube",
	"ujp":  "YouGlish JP",
	"jj":   "Jisho JP",
	"wiki": "Wikipedia",
	"ddg":  "DuckDuckGo",
}
var engine_uri = map[string]string{
	"wfr":  "https://en.wiktionary.org/wiki/{}#French",
	"fr":   "https://www.deepl.com/en/translator#en/fr/{}",
	"ufr":  "https://youglish.com/pronounce/{}/french/fr",
	"fren": "https://www.deepl.com/en/translator#fr/en/{}",
	"wen":  "https://en.wiktionary.org/wiki/{}#English",
	"uen":  "https://youglish.com/pronounce/{}/english/us",
	"wpt":  "https://en.wiktionary.org/wiki/{}#Portuguese",
	"sp":   "https://www.deepl.com/en/translator#en/es/{}",
	"spen": "https://www.deepl.com/en/translator#es/en/{}",
	"wsp":  "https://en.wiktionary.org/wiki/{}#Spanish",
	// "pt":   "https://www.deepl.com/en/translator#en/pt/{}",
	"wru":  "https://en.wiktionary.org/wiki/{}#Russian",
	"ru":   "https://www.deepl.com/en/translator#en/ru/{}",
	"uru":  "https://youglish.com/pronounce/{}/russian/ru",
	"ruen": "https://www.deepl.com/en/translator#ru/en/{}",
	"wnl":  "https://en.wiktionary.org/wiki/{}#Dutch",
	"nl":   "https://www.deepl.com/en/translator#en/nl/{}",
	"unl":  "https://youglish.com/pronounce/{}/dutch/nl",
	"nlen": "https://www.deepl.com/en/translator#nl/en/{}",
	"wge":  "https://en.wiktionary.org/wiki/{}#Dutch",
	"ge":   "https://www.deepl.com/en/translator#en/ge/{}",
	"uge":  "https://youglish.com/pronounce/{}/german/ge",
	"geen": "https://www.deepl.com/en/translator#ge/en/{}",
	"lt":   "https://www.deepl.com/en/translator#en/lt/{}",
	"lten": "https://www.deepl.com/en/translator#lt/en/{}",
	"wlt":  "https://en.wiktionary.org/wiki/{}#Lithuanian",
	"yt":   "https://www.youtube.com/results?search_query={}",
	"ujp":  "https://youglish.com/pronounce/{}/japanese",
	"jj":   "https://jisho.org/search/{}",
	"wiki": "https://en.wikipedia.org/wiki/{}",
	"ddg":  "https://duckduckgo.com/?t=ffab&q={}&ia=web",
}

func main() {
	//Ask for engine
	engine_used_str := run_rofi(true)
	engine_key := strings.Split(engine_used_str, "\t")[0]
	fmt.Println(engine_key)

	//Ask for prompt
	prompt := run_rofi(false)
	if !strings.HasSuffix(engine_key, "ge") && engine_key != "geen" {
		prompt = strings.ToLower(prompt)
	}
	fmt.Println(prompt)

	//Launch engine w/ prompt
	run_engine(engine_uri[engine_key], prompt)

}

func run_rofi(ask_for_engine bool) string {

	cmd := exec.Command("rofi")
	cmd.Args = append(cmd.Args, "-dmenu")
	cmd.Args = append(cmd.Args, "-i")
	cmd.Args = append(cmd.Args, "-font")
	cmd.Args = append(cmd.Args, "Roboto 14")
	cmd.Args = append(cmd.Args, "-l")
	cmd.Args = append(cmd.Args, "5")
	cmd.Args = append(cmd.Args, "-theme-str")
	cmd.Args = append(cmd.Args, "window { width: 40ch; }")

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
