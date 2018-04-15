// Code generated by reactGen. DO NOT EDIT.

package react

// PProps are the props for a <div> component
type PProps struct {
	AriaSet
	ClassName               string
	DangerouslySetInnerHTML *DangerousInnerHTML
	DataSet
	ID  string
	Key string

	OnChange
	OnClick

	Ref
	Role  string
	Style *CSS
}

func (p *PProps) assign(_v *_PProps) {

	if p.AriaSet != nil {
		for dk, dv := range p.AriaSet {
			_v.o.Set("aria-"+dk, dv)
		}
	}

	_v.ClassName = p.ClassName

	_v.DangerouslySetInnerHTML = p.DangerouslySetInnerHTML

	if p.DataSet != nil {
		for dk, dv := range p.DataSet {
			_v.o.Set("data-"+dk, dv)
		}
	}

	if p.ID != "" {
		_v.ID = p.ID
	}

	if p.Key != "" {
		_v.Key = p.Key
	}

	if p.OnChange != nil {
		_v.o.Set("onChange", p.OnChange.OnChange)
	}

	if p.OnClick != nil {
		_v.o.Set("onClick", p.OnClick.OnClick)
	}

	if p.Ref != nil {
		_v.o.Set("ref", p.Ref.Ref)
	}

	_v.Role = p.Role

	// TODO: until we have a resolution on
	// https://github.com/gopherjs/gopherjs/issues/236
	_v.Style = p.Style.hack()

}
