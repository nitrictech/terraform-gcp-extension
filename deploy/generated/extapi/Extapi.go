package extapi

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/nitrictech/terraform-gcp-extension/deploy/generated/extapi/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/nitrictech/terraform-gcp-extension/deploy/generated/extapi/internal"
)

// Defines an Extapi based on a Terraform module.
//
// Source at ./.ext/modules/extapi
type Extapi interface {
	cdktf.TerraformModule
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// The tree node.
	Node() constructs.Node
	// Experimental.
	Providers() *[]interface{}
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	SkipAssetCreationFromLocalModules() *bool
	// Experimental.
	Source() *string
	// Experimental.
	Version() *string
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	AddProvider(provider interface{})
	// Experimental.
	GetString(output *string) *string
	// Experimental.
	InterpolationForOutput(moduleOutput *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for Extapi
type jsiiProxy_Extapi struct {
	internal.Type__cdktfTerraformModule
}

func (j *jsiiProxy_Extapi) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Extapi) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Extapi) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Extapi) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Extapi) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Extapi) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Extapi) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Extapi) Providers() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"providers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Extapi) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Extapi) SkipAssetCreationFromLocalModules() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"skipAssetCreationFromLocalModules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Extapi) Source() *string {
	var returns *string
	_jsii_.Get(
		j,
		"source",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Extapi) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}


func NewExtapi(scope constructs.Construct, id *string, config *ExtapiConfig) Extapi {
	_init_.Initialize()

	if err := validateNewExtapiParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_Extapi{}

	_jsii_.Create(
		"extapi.Extapi",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

func NewExtapi_Override(e Extapi, scope constructs.Construct, id *string, config *ExtapiConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"extapi.Extapi",
		[]interface{}{scope, id, config},
		e,
	)
}

func (j *jsiiProxy_Extapi)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Extapi)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func Extapi_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateExtapi_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"extapi.Extapi",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Extapi_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateExtapi_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"extapi.Extapi",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Extapi) AddOverride(path *string, value interface{}) {
	if err := e.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (e *jsiiProxy_Extapi) AddProvider(provider interface{}) {
	if err := e.validateAddProviderParameters(provider); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addProvider",
		[]interface{}{provider},
	)
}

func (e *jsiiProxy_Extapi) GetString(output *string) *string {
	if err := e.validateGetStringParameters(output); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getString",
		[]interface{}{output},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Extapi) InterpolationForOutput(moduleOutput *string) cdktf.IResolvable {
	if err := e.validateInterpolationForOutputParameters(moduleOutput); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForOutput",
		[]interface{}{moduleOutput},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Extapi) OverrideLogicalId(newLogicalId *string) {
	if err := e.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (e *jsiiProxy_Extapi) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		e,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Extapi) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Extapi) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Extapi) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Extapi) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Extapi) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Extapi) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

