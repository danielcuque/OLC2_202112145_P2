package compiler

type Context struct {
	StackLabel []*Label
}

func NewContext() *Context {
	return &Context{
		StackLabel: make([]*Label, 0),
	}
}

func (c *Context) Push(label *Label) {
	c.StackLabel = append(c.StackLabel, label)
}

func (c *Context) Pop() *Label {
	label := c.StackLabel[len(c.StackLabel)-1]
	c.StackLabel = c.StackLabel[:len(c.StackLabel)-1]
	return label
}

func (c *Context) Peek() *Label {
	return c.StackLabel[len(c.StackLabel)-1]
}

func (c *Context) String() string {
	return ""
}
