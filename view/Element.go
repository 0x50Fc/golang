package view

type IElement interface {
	Id() int64
	Name() string
	Get(key string) interface{}
	Set(key string, value interface{})
	Attributes() map[string]interface{}
	FirstChild() IElement
	LastChild() IElement
	NextSibling() IElement
	PrevSibling() IElement
	Append(e IElement)
	Before(e IElement)
	After(e IElement)
	Parent() IElement
	Remove()

	setFirstChild(v IElement)
	setLastChild(v IElement)
	setNextSibling(v IElement)
	setPrevSibling(v IElement)
	setParent(v IElement)
}

type IElementConstructor = func(id int64, name string) IElement

var elementConstructorSet = map[string]IElementConstructor{}

func AddElementConstructor(name string, constructor IElementConstructor) {
	elementConstructorSet[name] = constructor
}

func NewElement(id int64, name string) IElement {
	fn, ok := elementConstructorSet[name]
	if ok {
		return fn(id, name)
	} else {
		return &Element{
			id:   id,
			name: name,
		}
	}
}

type Element struct {
	id         int64
	name       string
	attributes map[string]interface{}

	firstChild  IElement
	lastChild   IElement
	nextSibling IElement
	prevSibling IElement
	parent      IElement
}

func (E *Element) Id() int64 {
	return E.id
}

func (E *Element) SetId(v int64) {
	E.id = v
}

func (E *Element) Name() string {
	return E.name
}

func (E *Element) SetName(v string) {
	E.name = v
}

func (E *Element) Get(key string) interface{} {
	if E.attributes == nil {
		return nil
	}
	return E.attributes[key]
}

func (E *Element) Set(key string, v interface{}) {
	if v == nil {
		if E.attributes != nil {
			delete(E.attributes, key)
		}
	} else {
		if E.attributes == nil {
			E.attributes = map[string]interface{}{}
		}
		E.attributes[key] = v
	}
}

func (E *Element) Attributes() map[string]interface{} {
	if E.attributes == nil {
		E.attributes = map[string]interface{}{}
	}
	return E.attributes
}

func (E *Element) FirstChild() IElement {
	return E.firstChild
}

func (E *Element) LastChild() IElement {
	return E.lastChild
}

func (E *Element) NextSibling() IElement {
	return E.nextSibling
}

func (E *Element) PrevSibling() IElement {
	return E.prevSibling
}

func (E *Element) setFirstChild(v IElement) {
	E.firstChild = v
}

func (E *Element) setLastChild(v IElement) {
	E.lastChild = v
}

func (E *Element) setNextSibling(v IElement) {
	E.nextSibling = v
}

func (E *Element) setPrevSibling(v IElement) {
	E.prevSibling = v
}

func (E *Element) setParent(v IElement) {
	E.parent = v
}

func (E *Element) Append(e IElement) {

	e.Remove()

	if E.lastChild != nil {
		E.lastChild.setNextSibling(e)
		e.setPrevSibling(E.lastChild)
		E.lastChild = e
		e.setParent(E)
	} else {
		E.firstChild = e
		E.lastChild = e
		e.setParent(E)
	}

}

func (E *Element) Before(e IElement) {

	e.Remove()

	if E.prevSibling != nil {
		E.prevSibling.setNextSibling(e)
		e.setPrevSibling(E.prevSibling)
		e.setNextSibling(E)
		e.setParent(E.parent)
		E.prevSibling = e
	} else if E.parent != nil {
		e.setNextSibling(E)
		e.setParent(E.parent)
		E.setPrevSibling(e)
		E.parent.setFirstChild(e)
	}

}

func (E *Element) After(e IElement) {

	e.Remove()

	if E.nextSibling != nil {
		E.nextSibling.setPrevSibling(e)
		e.setNextSibling(E.nextSibling)
		e.setPrevSibling(E)
		e.setParent(E.parent)
		E.setNextSibling(e)
	} else if E.parent != nil {
		e.setPrevSibling(E)
		e.setParent(E.parent)
		E.nextSibling = e
		E.parent.setLastChild(e)
	}
}

func (E *Element) Remove() {

	if E.prevSibling != nil && E.parent != nil {
		E.prevSibling.setNextSibling(E.nextSibling)
		if E.nextSibling != nil {
			E.nextSibling.setPrevSibling(E.prevSibling)
		} else {
			E.parent.setLastChild(E.prevSibling)
		}
	} else if E.parent != nil {
		E.parent.setFirstChild(E.nextSibling)
		if E.nextSibling != nil {
			E.nextSibling.setPrevSibling(nil)
		} else {
			E.parent.setLastChild(nil)
		}
	}

	E.parent = nil
	E.prevSibling = nil
	E.nextSibling = nil
}

func (E *Element) Parent() IElement {
	return E.parent
}
