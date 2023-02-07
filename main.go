package main

import (
	"archive/zip"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"reflect"
	"strings"
	"text/template"

	"github.com/tinode/jsonco"
)

// Theme ...
type Theme struct {
	Name   string
	URL    string
	Target string
	File   string
	Dark   bool
	HC     bool
}

type ColorDefaults struct {
	light   string
	dark    string
	hcLight string
	hcDark  string
}

func main() {
	themes := []Theme{
		{
			Name:   "dracula",
			URL:    "https://marketplace.visualstudio.com/_apis/public/gallery/publishers/dracula-theme/vsextensions/theme-dracula/2.24.2/vspackage",
			Target: "dracula",
			File:   "extension/theme/dracula.json",
			Dark:   true,
		},
		{
			Name:   "solarized-light",
			URL:    "https://marketplace.visualstudio.com/_apis/public/gallery/publishers/ryanolsonx/vsextensions/solarized/2.0.3/vspackage",
			Target: "solarized",
			File:   "extension/themes/light-color-theme.json",
			Dark:   false,
		},
		{
			Name:   "solarized-dark",
			URL:    "https://marketplace.visualstudio.com/_apis/public/gallery/publishers/ryanolsonx/vsextensions/solarized/2.0.3/vspackage",
			Target: "solarized",
			File:   "extension/themes/dark-color-theme.json",
			Dark:   true,
		},
		{
			Name: "material-light",

			URL:    "https://marketplace.visualstudio.com/_apis/public/gallery/publishers/Equinusocio/vsextensions/vsc-material-theme/33.4.0/vspackage",
			Target: "material",
			File:   "extension/build/themes/Material-Theme-Lighter.json",
			Dark:   false,
		},
		{
			Name:   "material-dark",
			URL:    "https://marketplace.visualstudio.com/_apis/public/gallery/publishers/Equinusocio/vsextensions/vsc-material-theme/33.4.0/vspackage",
			Target: "material",
			File:   "extension/build/themes/Material-Theme-Default.json",
			Dark:   true,
		},
		{
			Name:   "github-light",
			URL:    "https://marketplace.visualstudio.com/_apis/public/gallery/publishers/GitHub/vsextensions/github-vscode-theme/6.0.0/vspackage",
			Target: "github",
			File:   "extension/themes/light.json",
			Dark:   false,
		},
		{
			Name:   "github-dark",
			URL:    "https://marketplace.visualstudio.com/_apis/public/gallery/publishers/GitHub/vsextensions/github-vscode-theme/6.0.0/vspackage",
			Target: "github",
			File:   "extension/themes/dark.json",
			Dark:   true,
		},
		{
			Name:   "aura",
			URL:    "https://marketplace.visualstudio.com/_apis/public/gallery/publishers/DaltonMenezes/vsextensions/aura-theme/2.1.2/vspackage",
			Target: "aura",
			File:   "extension/themes/aura-soft-dark-color-theme.json",
			Dark:   true,
		},
		{
			Name:   "tokyo-night",
			URL:    "https://marketplace.visualstudio.com/_apis/public/gallery/publishers/enkia/vsextensions/tokyo-night/0.9.4/vspackage",
			Target: "tokyo-night",
			File:   "extension/themes/tokyo-night-color-theme.json",
			Dark:   true,
		},
		{
			Name:   "tokyo-night-storm",
			URL:    "https://marketplace.visualstudio.com/_apis/public/gallery/publishers/enkia/vsextensions/tokyo-night/0.9.4/vspackage",
			Target: "tokyo-night-storm",
			File:   "extension/themes/tokyo-night-storm-color-theme.json",
			Dark:   true,
		},
		// vscode light themes
		{
			Name:   "light",
			URL:    "https://raw.githubusercontent.com/microsoft/vscode/main/extensions/theme-defaults/themes/light_vs.json",
			Target: "light",
			Dark:   false,
		},
		{
			Name:   "light-plus",
			URL:    "https://raw.githubusercontent.com/microsoft/vscode/main/extensions/theme-defaults/themes/light_plus.json",
			Target: "light-plus",
			Dark:   false,
		},
		{
			Name:   "light-plus-v2",
			URL:    "https://raw.githubusercontent.com/microsoft/vscode/main/extensions/theme-defaults/themes/light_plus_experimental.json",
			Target: "light-plus-v2",
			Dark:   false,
		},
		{
			Name:   "quiet-light",
			URL:    "https://raw.githubusercontent.com/microsoft/vscode/main/extensions/theme-quietlight/themes/quietlight-color-theme.json",
			Target: "quiet-light",
			Dark:   false,
		},
		{
			Name:   "solarized-light",
			URL:    "https://raw.githubusercontent.com/microsoft/vscode/main/extensions/theme-solarized-light/themes/solarized-light-color-theme.json",
			Target: "solarized-light",
			Dark:   false,
		},
		// vscode dark themes
		{
			Name:   "abyss",
			URL:    "https://raw.githubusercontent.com/microsoft/vscode/main/extensions/theme-abyss/themes/abyss-color-theme.json",
			Target: "abyss",
			Dark:   true,
		},
		{
			Name:   "dark",
			URL:    "https://raw.githubusercontent.com/microsoft/vscode/main/extensions/theme-defaults/themes/dark_vs.json",
			Target: "dark",
			Dark:   true,
		},
		{
			Name:   "dark-plus",
			URL:    "https://raw.githubusercontent.com/microsoft/vscode/main/extensions/theme-defaults/themes/dark_plus.json",
			Target: "dark-plus",
			Dark:   true,
		},
		{
			Name:   "dark-plus-v2",
			URL:    "https://raw.githubusercontent.com/microsoft/vscode/main/extensions/theme-defaults/themes/dark_plus_experimental.json",
			Target: "dark-plus-v2",
			Dark:   true,
		},
		{
			Name:   "kimbie-dark",
			URL:    "https://raw.githubusercontent.com/microsoft/vscode/main/extensions/theme-kimbie-dark/themes/kimbie-dark-color-theme.json",
			Target: "kimbie-dark",
			Dark:   true,
		},
		{
			Name:   "monokai",
			URL:    "https://raw.githubusercontent.com/microsoft/vscode/main/extensions/theme-monokai/themes/monokai-color-theme.json",
			Target: "monokai",
			Dark:   true,
		},
		{
			Name:   "monokai-dimmed",
			URL:    "https://raw.githubusercontent.com/microsoft/vscode/main/extensions/theme-monokai-dimmed/themes/dimmed-monokai-color-theme.json",
			Target: "monokai-dimmed",
			Dark:   true,
		},
		{
			Name:   "red",
			URL:    "https://raw.githubusercontent.com/microsoft/vscode/main/extensions/theme-red/themes/Red-color-theme.json",
			Target: "red",
			Dark:   true,
		},
		{
			Name:   "solarized-dark",
			URL:    "https://raw.githubusercontent.com/microsoft/vscode/main/extensions/theme-solarized-dark/themes/solarized-dark-color-theme.json",
			Target: "solarized-dark",
			Dark:   true,
		},
		{
			Name:   "tomorrow-night-blue",
			URL:    "https://raw.githubusercontent.com/microsoft/vscode/main/extensions/theme-tomorrow-night-blue/themes/tomorrow-night-blue-color-theme.json",
			Target: "tomorrow-night-blue",
			Dark:   true,
		},
		// vscode high contrast themes
		{
			Name:   "high-contrast-light",
			URL:    "https://raw.githubusercontent.com/microsoft/vscode/main/extensions/theme-defaults/themes/hc_light.json",
			Target: "high-contrast-light",
			Dark:   false,
			HC:     true,
		},
		{
			Name:   "high-contrast-dark",
			URL:    "https://raw.githubusercontent.com/microsoft/vscode/main/extensions/theme-defaults/themes/hc_black.json",
			Target: "high-contrast-dark",
			Dark:   true,
			HC:     true,
		},
	}

	for _, theme := range themes {
		fmt.Println("Process theme: ", theme.Name)

		var err error
		var content []byte

		if _, err = os.Stat("./tmp/" + theme.Target + ".zip"); os.IsNotExist(err) {
			fmt.Println("  Download theme")
			downloadTheme(theme)
		}

		if theme.File != "" {
			fmt.Println("  Extract theme")
			content, err = extractTheme(theme)
		} else {
			content, err = extractThemeJson(theme)
		}

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("  Generate template")
		generateTheme(theme, content)
	}
}

func generateTheme(theme Theme, content []byte) {
	params := makeTemplateParams(theme, content)

	t, err := template.ParseFiles("./template.js")
	if err != nil {
		log.Fatal(err)
	}

	out, err := os.Create("./theme/" + theme.Name + ".ts")
	if err != nil {
		log.Fatal(err)
	}

	err = t.Execute(out, params)
	if err != nil {
		log.Fatal(err)
	}

	out.Close()
}

// Style ...
type Style struct {
	Color     *string
	FontStyle *string
	Prio      *int
}

// TokenColorSettings ...
type TokenColorSettings struct {
	Foreground *string
	FontStyle  *string
}

// TokenColor ...
type TokenColor struct {
	Scope    interface{}
	Settings TokenColorSettings
}

// VsCodeTheme ...
type VsCodeTheme struct {
	Colors      map[string]string
	TokenColors []TokenColor
}

func find(data VsCodeTheme, theme Theme, defaults ColorDefaults, keys ...string) Style {
	style := Style{}

	for _, key := range keys {
		if value, exist := data.Colors[key]; exist && style.Color == nil {
			return Style{Color: &value}
		}

		for _, tokenColor := range data.TokenColors {
			scopes := []string{}
			rt := reflect.TypeOf(tokenColor.Scope)
			if tokenColor.Scope == nil {
				continue
			}

			switch rt.Kind() {
			case reflect.Slice:
				for _, s := range tokenColor.Scope.([]interface{}) {
					scopes = append(scopes, s.(string))
				}
			case reflect.String:
				splitted := strings.Split(tokenColor.Scope.(string), ",")
				for _, s := range splitted {
					scopes = append(scopes, strings.TrimSpace(s))
				}
			default:
				panic(fmt.Sprintf("Unecpected scope type %s", rt))
			}

			for i, scope := range scopes {
				if scope == key &&
					(style.Color == nil || *style.Prio > i) &&
					tokenColor.Settings.Foreground != nil {
					prio := i
					style.Color = tokenColor.Settings.Foreground
					style.Prio = &prio
				}

				if scope == key && style.FontStyle == nil && tokenColor.Settings.FontStyle != nil {
					style.FontStyle = tokenColor.Settings.FontStyle
				}
			}
		}
	}

	if style.Color == nil {
		if theme.Dark {
			if theme.HC {
				return Style{Color: &defaults.hcDark}
			} else {
				return Style{Color: &defaults.dark}
			}
		} else {
			if theme.HC {
				return Style{Color: &defaults.hcLight}
			} else {
				return Style{Color: &defaults.light}
			}
		}
		panic(fmt.Sprintf("Could not find color by: %s", keys))
	}

	return style
}

// TemplateParams ...
type TemplateParams struct {
	ExportPrefix string
	Dark         bool

	// Editor
	Background         Style
	Foreground         Style
	Selection          Style
	Cursor             Style
	DropdownBackground Style
	DropdownBorder     Style
	ActiveLine         Style
	MatchingBracket    Style

	// Syntax
	Keyword   Style // if else, etc
	Storage   Style // const, let, etc - Not supported in CM
	Parameter Style // fn(parmater)    - Not supported in CM
	Variable  Style
	Function  Style
	String    Style
	Constant  Style // ???
	Type      Style // x: MyType
	Class     Style // class MyClass
	Number    Style
	Comment   Style
	Heading   Style
	Invalid   Style
	Regexp    Style
}

func makeTemplateParams(theme Theme, content []byte) TemplateParams {
	var data VsCodeTheme
	jr := jsonco.New(bytes.NewBuffer(content))
	json.NewDecoder(jr).Decode(&data)

	params := TemplateParams{
		ExportPrefix: kebabToCamelCase(theme.Name),
		Dark:         theme.Dark,
		// Layout
		Background:         find(data, theme, ColorDefaults{light: "#ffffff", dark: "#1E1E1E", hcDark: "#000000", hcLight: "#ffffff"}, "editor.background"),
		Foreground:         find(data, theme, ColorDefaults{light: "#333333", dark: "#BBBBBB", hcDark: "#ffffff", hcLight: "#292929"}, "foreground", "input.foreground", "editor.foreground"),
		Selection:          find(data, theme, ColorDefaults{light: "#ADD6FF", dark: "#264F78", hcDark: "#f3f518", hcLight: "#0F4A85"}, "editor.selectionBackground"),
		Cursor:             find(data, theme, ColorDefaults{light: "#333333", dark: "#BBBBBB", hcDark: "#ffffff", hcLight: "#292929"}, "editorCursor.foreground", "foreground"),
		DropdownBackground: find(data, theme, ColorDefaults{dark: "#252526", light: "#F3F3F3", hcDark: "#0C141F", hcLight: "#ffffff"}, "editor.background"),
		DropdownBorder:     find(data, theme, ColorDefaults{dark: "#454545", light: "#C8C8C8", hcDark: "#6FC3DF", hcLight: "#0F4A85"}, "dropdown.border", "foreground", "editorSuggestWidget.border"),
		ActiveLine:         find(data, theme, ColorDefaults{light: "#ADD6FF", dark: "#264F78", hcDark: "#f3f518", hcLight: "#0F4A85"}, "editor.lineHighlightBackground", "editor.selectionBackground"),
		MatchingBracket:    find(data, theme, ColorDefaults{dark: "#0064001a", light: "#0064001a", hcDark: "#0064001a", hcLight: "#0000"}, "editorBracketMatch.background", "editor.lineHighlightBackground", "editor.selectionBackground"),
		// Syntax
		Keyword:   find(data, theme, ColorDefaults{light: "#0000ff", dark: "#569cd6"}, "keyword"),
		Storage:   find(data, theme, ColorDefaults{light: "#0000ff", dark: "#569cd6"}, "storage", "keyword"),
		Variable:  find(data, theme, ColorDefaults{light: "#0070c1", dark: "#4fc1ff"}, "variable", "variable.parameter", "variable.other", "variable.language", "foreground"),
		Parameter: find(data, theme, ColorDefaults{light: "#333333", dark: "#BBBBBB"}, "variable.parameter", "variable.other", "variable"),
		Function:  find(data, theme, ColorDefaults{light: "#795e26", dark: "#dcdcaa"}, "entity.name.function", "support.function", "entity.name", "support"),
		String:    find(data, theme, ColorDefaults{light: "#a31515", dark: "#ce9178"}, "string"),
		Constant:  find(data, theme, ColorDefaults{light: "#333333", dark: "#BBBBBB"}, "constant", "constant.character", "constant.keyword"),
		Type:      find(data, theme, ColorDefaults{light: "#267f99", dark: "#4ec9b0"}, "entity.name.type", "entity.name.class", "support.type", "support"),
		Class:     find(data, theme, ColorDefaults{light: "#267f99", dark: "#4ec9b0"}, "entity.name.class", "entity.name", "entity"),
		Number:    find(data, theme, ColorDefaults{light: "#098658", dark: "#b5cea8"}, "constant.numeric", "constant"),
		Comment:   find(data, theme, ColorDefaults{light: "#008000", dark: "#6a9955"}, "comment"),
		Heading:   find(data, theme, ColorDefaults{light: "#000080", dark: "#000080"}, "markup.heading", "markup.heading.setext", "heading.1.markdown entity.name"),
		Invalid:   find(data, theme, ColorDefaults{light: "#cd3131", dark: "#f44747"}, "invalid", "editorError.foreground", "errorForeground", "foreground", "input.foreground"),
		Regexp:    find(data, theme, ColorDefaults{light: "#811f3f", dark: "#646695"}, "string.regexp", "string"),
	}

	return params
}

func extractTheme(theme Theme) ([]byte, error) {
	r, err := zip.OpenReader("tmp/" + theme.Target + ".zip")

	if err != nil {
		log.Fatal(err)
	}

	defer r.Close()

	for _, f := range r.File {
		if f.Name != theme.File {
			continue
		}

		rc, err := f.Open()
		if err != nil {
			log.Fatal(err)
		}

		content, err := ioutil.ReadAll(rc)
		if err != nil {
			log.Fatal(err)
		}

		rc.Close()

		return content, nil
	}

	return nil, fmt.Errorf("Cound not find file %s in extension", theme.File)
}

func extractThemeJson(theme Theme) ([]byte, error) {
	file, err := os.Open("tmp/" + theme.Target + ".zip")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	content, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	return content, nil
}

func downloadTheme(theme Theme) {

	resp, err := http.Get(theme.URL)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Fatal("Could not download theme: ", theme, "StatusCode: ", resp.StatusCode)
	}

	_ = os.Mkdir("./tmp", 0700)

	// Create the file
	out, err := os.Create("tmp/" + theme.Target + ".zip")
	if err != nil {
		log.Fatal(err)
	}

	defer out.Close()

	_, err = io.Copy(out, resp.Body)

	if err != nil {
		log.Fatal(err)
	}
}

func kebabToCamelCase(kebab string) (camelCase string) {
	isToUpper := false
	for _, runeValue := range kebab {
		if isToUpper {
			camelCase += strings.ToUpper(string(runeValue))
			isToUpper = false
		} else {
			if runeValue == '-' {
				isToUpper = true
			} else {
				camelCase += string(runeValue)
			}
		}
	}
	return
}
