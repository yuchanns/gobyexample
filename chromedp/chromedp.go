package chromedp

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/chromedp/chromedp"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

type Mathematic struct {
	ID     string
	Expr   string
	Inline bool
}

type Mathmatics []*Mathematic

func (m Mathmatics) String() string {
	paragraphs := make([]string, 0, len(m))
	for _, mathmatic := range m {
		delimiter := "$$"
		if mathmatic.Inline {
			delimiter = "$"
		}
		paragraphs = append(paragraphs, fmt.Sprintf(
			"<p id=\"%s\">%s%s%s</p>",
			mathmatic.ID,
			delimiter,
			mathmatic.Expr,
			delimiter,
		))
	}
	return strings.Join(paragraphs, "")
}

func (m Mathmatics) GenerateActions() []chromedp.Action {
	actions := make([]chromedp.Action, 0, len(m))
	for _, mathmatic := range m {
		actions = append(actions, chromedp.InnerHTML(
			fmt.Sprintf("#%s", mathmatic.ID),
			&mathmatic.Expr,
			chromedp.NodeVisible,
			chromedp.ByQuery,
		))
	}
	return actions
}

func (m Mathmatics) Clone() Mathmatics {
	jsonByte, _ := json.Marshal(m)
	var cpy Mathmatics
	_ = json.Unmarshal(jsonByte, &cpy)
	return cpy
}

func (m Mathmatics) GenerateTempURL() (string, func() error, error) {
	content := fmt.Sprintf(`<!DOCTYPE html><html><head><script>MathJax={tex:{inlineMath:[['$','$']]},svg:{fontCache:'global'}};</script><script id="MathJax-script"async src="https://cdn.jsdelivr.net/npm/mathjax@3/es5/tex-svg.js"></script></head><body>%s</body></html>`, m)

	tempFile, err := ioutil.TempFile("", "*.html")
	if err != nil {
		return "", nil, err
	}
	if _, err := tempFile.WriteString(content); err != nil {
		return "", nil, err
	}
	return fmt.Sprintf("file:///%s", tempFile.Name()), func() error {
		return os.Remove(tempFile.Name())
	}, nil
}

func Render(ctx context.Context, ms Mathmatics) (Mathmatics, string, string, error) {
	mathematics := ms.Clone()
	url, tmpClose, err := mathematics.GenerateTempURL()
	if err != nil {
		return nil, "", "", err
	}
	defer tmpClose()
	chromeCtx, chromeCancel := chromedp.NewContext(ctx, chromedp.WithLogf(log.Printf))
	defer chromeCancel()
	var (
		html string
		svg  string
	)
	actions := []chromedp.Action{
		chromedp.Navigate(url),
		chromedp.WaitVisible(".MathJax"),
		chromedp.InnerHTML("html", &html, chromedp.NodeVisible,
			chromedp.ByQuery),
		chromedp.OuterHTML("#MJX-SVG-global-cache", &svg,
			chromedp.ByQuery),
	}
	actions = append(actions, mathematics.GenerateActions()...)

	if err := chromedp.Run(chromeCtx, actions...); err != nil {
		return nil, "", "", err
	}
	reg := regexp.MustCompile("<style(([\\s\\S])*?)</style>")
	matches := reg.FindAllStringSubmatch(html, -1)
	matchJoins := make([]string, 0, len(matches))
	for _, match := range matches {
		matchJoins = append(matchJoins, match[0])
	}
	return mathematics, strings.Join(matchJoins, ""), svg, nil
}
