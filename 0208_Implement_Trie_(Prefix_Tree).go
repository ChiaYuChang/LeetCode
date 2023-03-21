package leetcode

type Q0208 struct{}

type Q0208Trie struct {
	root *Q0208Node
}

func (q Q0208) NewQ0208Trie() Q0208Trie {
	return Q0208Trie{root: &Q0208Node{children: make([]*Q0208Node, 26)}}
}

func (this *Q0208Trie) Insert(word string) {
	curr := this.root
	for i := 0; i < len(word); i++ {
		w := word[i]
		if !curr.HasChild(w) {
			curr.AppendNode(w)
		}
		curr = curr.GetChild(w)
	}
	curr.isEnd = true
}

func (this *Q0208Trie) SearchNode(word string) *Q0208Node {
	curr := this.root
	for i := 0; i < len(word); i++ {
		w := word[i]
		if !curr.HasChild(w) {
			return nil
		}
		curr = curr.GetChild(w)
	}
	return curr
}

func (this *Q0208Trie) StartsWith(prefix string) bool {
	return this.SearchNode(prefix) != nil
}

func (this *Q0208Trie) Search(word string) bool {
	n := this.SearchNode(word)
	if n == nil {
		return false
	}
	return n.isEnd
}

type Q0208Node struct {
	isEnd    bool
	children []*Q0208Node
}

func (n *Q0208Node) HasChild(char byte) bool {
	return n.children[char-'a'] != nil
}

func (n *Q0208Node) GetChild(char byte) *Q0208Node {
	return n.children[char-'a']
}

func (n *Q0208Node) AppendNode(char byte) {
	if n.HasChild(char) {
		return
	}
	n.children[char-'a'] = &Q0208Node{children: make([]*Q0208Node, 26)}
}
