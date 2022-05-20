package main

import "encoding/json"

/*Policies interface is use to store and retrive different kind
of policies*/
type Policies interface {
	//returns the kind of policy
	GetKind() string

	//returns the ApiVesion of policy
	GetApiVersion() string

	//returns the MetaData of policy
	GetMetaData() map[string]string

	//returns the selector of policy
	GetSpecSelector() Selector

	//returns the spec data in slice of bytes
	GetSpecData() []byte
}

//Implemention Of Polices for Normal Policy Struct
func (p Policy) GetKind() string { return p.Kind }

func (p Policy) GetApiVersion() string { return p.APIVersion }

func (p Policy) GetMetaData() map[string]string { return p.Metadata }

func (p Policy) GetSpecSelector() Selector { return p.Spec.Selector }

//func (p Policy) Getspec() {} { return p.Spec }

func (p Policy) GetSpecData() []byte {

	data, _ := json.Marshal(&p.Spec)
	return data
}
