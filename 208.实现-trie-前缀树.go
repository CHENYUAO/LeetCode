/*
 * @lc app=leetcode.cn id=208 lang=golang
 *
 * [208] 实现 Trie (前缀树)
 */

// @lc code=start
type Trie struct {
	isWord bool
	childs [26]*Trie
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	w_rune := []rune(word)
	curr := this
	for _, item := range w_rune {
		if curr.childs[item-'a'] == nil {
			// 建节点
			curr.childs[item-'a'] = &Trie{}
		}
		curr = curr.childs[item-'a']
	}
	curr.isWord = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	w_rune := []rune(word)
	curr := this
	for _, item := range w_rune {
		if curr.childs[item-'a'] == nil {
			return false
		}
		curr = curr.childs[item-'a']
	}
	return curr.isWord
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	p_rune := []rune(prefix)
	curr := this
	for _, item := range p_rune {
		if curr.childs[item-'a'] == nil {
			return false
		}
		curr = curr.childs[item-'a']
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
// @lc code=end

