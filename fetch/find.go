/* fetch package implements those functions that
shouldn't be included in soup.go so as to not put it
within the user's reach
*/

package fetch

import "golang.org/x/net/html"

var (
	nodeLinks []*html.Node
	keyvalues map[string]string
)

func Set() {
	nodeLinks = make([]*html.Node, 0, 10)
	keyvalues = make(map[string]string)
}

// Using depth first search to find the first occurrence and return
func FindOnce(n *html.Node, args []string, uni bool) (*html.Node, bool, bool) {
	if uni == true {
		if n.Type == html.ElementNode && n.Data == args[0] {
			if len(args) > 1 && len(args) < 4 {
				for i := 0; i < len(n.Attr); i++ {
					if n.Attr[i].Key == args[1] && n.Attr[i].Val == args[2] {
						return n, true, true
					}
				}
			} else if len(args) == 1 {
				return n, true, true
			}
		}
	}
	uni = true
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		p, q, _ := FindOnce(c, args, true)
		if q != false {
			return p, q, true
		}
	}
	return nil, false, true
}

// Using depth first search to find all occurrences and return
func FindAllofem(n *html.Node, args []string, uni bool) ([]*html.Node, bool, bool) {
	if uni == true {
		if n.Data == args[0] {
			if len(args) > 1 && len(args) < 4 {
				for i := 0; i < len(n.Attr); i++ {
					if n.Attr[i].Key == args[1] && n.Attr[i].Val == args[2] {
						nodeLinks = append(nodeLinks, n)
					}
				}
			} else if len(args) == 1 {
				nodeLinks = append(nodeLinks, n)
			}
		}
	}
	uni = true
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		FindAllofem(c, args, true)
	}
	return nodeLinks, true, true
}

// Returns a key pair value (like a dictionary) for each attribute
func GetKeyValue(attributes []html.Attribute) map[string]string {
	for i := 0; i < len(attributes); i++ {
		_, exists := keyvalues[attributes[i].Key]
		if exists == false {
			keyvalues[attributes[i].Key] = attributes[i].Val
		}
	}
	return keyvalues
}
