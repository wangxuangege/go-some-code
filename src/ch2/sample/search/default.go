package search

type defaultMatcher struct{}

func init() {
	var mather defaultMatcher
	Register("default", mather)
}

func (m defaultMatcher) Search(feed *Feed, searchTerm string) ([]*Result, error) {
	return nil, nil
}
