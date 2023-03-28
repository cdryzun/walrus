package runtime

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"path"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/gin-gonic/gin/render"
	"k8s.io/apimachinery/pkg/util/sets"

	"github.com/seal-io/seal/pkg/apis/auth/session"
	"github.com/seal-io/seal/utils/log"
	"github.com/seal-io/seal/utils/strs"
)

func init() {
	// disable gin default binding
	binding.Validator = nil
}

// MustRoutePost likes RoutePost,
// but panics if error occurs.
func MustRoutePost(r gin.IRoutes, p string, h any) {
	var err = RoutePost(r, p, h)
	if err != nil {
		panic(err)
	}
}

// RoutePost registers the given handler as POST router.
func RoutePost(r gin.IRoutes, p string, h any) error {
	switch t := h.(type) {
	default:
		return errors.New("invalid handler")
	case ErrorHandler:
		r.POST(p, WrapErrorHandler(t))
	case Handler:
		r.POST(p, WrapHandler(t))
	case ErrorHandle:
		r.POST(p, WrapErrorHandle(t))
	case Handle:
		r.POST(p, t)
	case HTTPHandle:
		r.POST(p, WrapHTTPHandle(t))
	case HTTPHandler:
		r.POST(p, WrapHTTPHandler(t))
	}
	return nil
}

// MustRouteDelete likes RouteDelete,
// but panics if error occurs.
func MustRouteDelete(r gin.IRoutes, p string, h any) {
	var err = RouteDelete(r, p, h)
	if err != nil {
		panic(err)
	}
}

// RouteDelete registers the given handler as DELETE router.
func RouteDelete(r gin.IRoutes, p string, h any) error {
	switch t := h.(type) {
	default:
		return errors.New("invalid handler")
	case ErrorHandler:
		r.DELETE(p, WrapErrorHandler(t))
	case Handler:
		r.DELETE(p, WrapHandler(t))
	case ErrorHandle:
		r.DELETE(p, WrapErrorHandle(t))
	case Handle:
		r.DELETE(p, t)
	case HTTPHandle:
		r.DELETE(p, WrapHTTPHandle(t))
	case HTTPHandler:
		r.DELETE(p, WrapHTTPHandler(t))
	}
	return nil
}

// MustRoutePut likes RoutePut,
// but panics if error occurs.
func MustRoutePut(r gin.IRoutes, p string, h any) {
	var err = RoutePut(r, p, h)
	if err != nil {
		panic(err)
	}
}

// RoutePut registers the given handler as PUT router.
func RoutePut(r gin.IRoutes, p string, h any) error {
	switch t := h.(type) {
	default:
		return errors.New("invalid handler")
	case ErrorHandler:
		r.PUT(p, WrapErrorHandler(t))
	case Handler:
		r.PUT(p, WrapHandler(t))
	case ErrorHandle:
		r.PUT(p, WrapErrorHandle(t))
	case Handle:
		r.PUT(p, t)
	case HTTPHandle:
		r.PUT(p, WrapHTTPHandle(t))
	case HTTPHandler:
		r.PUT(p, WrapHTTPHandler(t))
	}
	return nil
}

// MustRouteGet likes RouteGet,
// but panics if error occurs.
func MustRouteGet(r gin.IRoutes, p string, h any) {
	var err = RouteGet(r, p, h)
	if err != nil {
		panic(err)
	}
}

// RouteGet registers the given handler as GET router.
func RouteGet(r gin.IRoutes, p string, h any) error {
	switch t := h.(type) {
	default:
		return errors.New("invalid handler")
	case ErrorHandler:
		r.GET(p, WrapErrorHandler(t))
	case Handler:
		r.GET(p, WrapHandler(t))
	case ErrorHandle:
		r.GET(p, WrapErrorHandle(t))
	case Handle:
		r.GET(p, t)
	case HTTPHandle:
		r.GET(p, WrapHTTPHandle(t))
	case HTTPHandler:
		r.GET(p, WrapHTTPHandler(t))
	}
	return nil
}

// MustRouteHead likes RouteHead,
// but panics if error occurs.
func MustRouteHead(r gin.IRoutes, p string, h any) {
	var err = RouteHead(r, p, h)
	if err != nil {
		panic(err)
	}
}

// RouteHead registers the given handler as HEAD router.
func RouteHead(r gin.IRoutes, p string, h any) error {
	switch t := h.(type) {
	default:
		return errors.New("invalid handler")
	case ErrorHandler:
		r.HEAD(p, WrapErrorHandler(t))
	case Handler:
		r.HEAD(p, WrapHandler(t))
	case ErrorHandle:
		r.HEAD(p, WrapErrorHandle(t))
	case Handle:
		r.HEAD(p, t)
	case HTTPHandle:
		r.HEAD(p, WrapHTTPHandle(t))
	case HTTPHandler:
		r.HEAD(p, WrapHTTPHandler(t))
	}
	return nil
}

// MustRouteStatic likes RouteStatic,
// but panics if error occurs.
func MustRouteStatic(r gin.IRoutes, p string, h any) {
	var err = RouteStatic(r, p, h)
	if err != nil {
		panic(err)
	}
}

// RouteStatic registers the given handler as HEAD and GET router.
func RouteStatic(r gin.IRoutes, p string, h any) error {
	var err = RouteHead(r, p, h)
	if err != nil {
		return err
	}
	err = RouteGet(r, p, h)
	if err != nil {
		return err
	}
	return nil
}

// MustRouteResource likes RouteResource,
// but panics if error occurs.
func MustRouteResource(r gin.IRoutes, h Resource) {
	var err = RouteResource(r, h)
	if err != nil {
		panic(err)
	}
}

// RouteResource reflects the function descriptors of the given handler,
// and registers as router if satisfy the below rules.
//
//	Input : struct type.
//	Output: any types.
//
//	* Basic APIs
//
//	func Stream(RequestStream, <Input>) error
//	 ->     ws /<plural>/:id
//	func Create(*gin.Context, <Input>) ((<Output>,) error)
//	 ->   POST /<plural>
//	func Delete(*gin.Context, <Input>) ((<Output>,) error)
//	 -> DELETE /<plural>/:id
//	func Update(*gin.Context, <Input>) ((<Output>,) error)
//	 ->    PUT /<plural>/:id
//	func Get(*gin.Context, <Input>) (<Output>, error)
//	 ->    GET /<plural>/:id
//
//	* Batch APIs
//
//	func CollectionStream(RequestStream, <Input>) error
//	 ->     ws /<plural>
//	func CollectionCreate(*gin.Context, <Input>) ((<Output>,) error)
//	 ->   POST /<plural>/_/batch
//	func CollectionDelete(*gin.Context, <Input>) ((<Output>,) error)
//	 -> DELETE /<plural>
//	func CollectionUpdate(*gin.Context, <Input>) ((<Output>,) error)
//	 ->    PUT /<plural>
//	func CollectionGet(*gin.Context, <Input>) (<Output>, (int,) error)
//	 ->    GET /<plural>
//
//	* Extensional APIs
//
//	func Stream<Something>(RequestStream, <Input>) error
//	 ->     ws /<plural>/:id/<something>
//	func CollectionStream<Something>(RequestStream, <Input>) error
//	 ->     ws /<plural>/:id/<something>
//	func Create<Something>(*gin.Context, <Input>) ((<Output>,) error)
//	 ->   POST /<plural>/:id/<something>
//	func CollectionCreate<Something>(*gin.Context, <Input>) ((<Output>,) error)
//	 ->   POST /<plural>/_/<something>
//	func Delete<Something>(*gin.Context, <Input>) ((<Output>,) error)
//	 -> DELETE /<plural>/:id/<something>
//	func CollectionDelete<Something>(*gin.Context, <Input>) ((<Output>,) error)
//	 -> DELETE /<plural>/_/<something>
//	func Update<Something>(*gin.Context, <Input>) ((<Output>,) error)
//	 ->    PUT /<plural>/:id/<something>
//	func CollectionUpdate<Something>(*gin.Context, <Input>) ((<Output>,) error)
//	 ->    PUT /<plural>/_/<something>
//	func Get<Something>(*gin.Context, <Input>) (<Output>, (int,) error)
//	 ->    GET /<plural>/:id/<something>
//	func CollectionGet<Something>(*gin.Context, <Input>) (<Output>, (int,) error)
//	 ->    GET /<plural>/_/<something>
//	func Route<Something>(*gin.Context, <Input(route:method=subpath)>) ((<Output>), (int,) error)
//	 -> method /<plural>/:id/<subpath>
//	func CollectionRoute<Something>(*gin.Context, <Input(route:method=subpath)>) ((<Output>), (int,) error)
//	 -> method /<plural>/_/<subpath>
func RouteResource(r gin.IRoutes, h Resource) error {
	if adv, ok := r.(AdviceBeforeResourceRegistering); ok {
		var err = adv.BeforeAdvice(adviceResource{Resource: h})
		if err != nil {
			return err
		}
	}

	var k = h.Kind()
	var resource, resourcePath = getResourceAndResourcePath(k)
	var rhs = getRouteHandlers(h, resourcePath)
	for i := range rhs {
		var rh = rhs[i]

		// normal request caller.
		var (
			mr      reflect.Value
			mrt     reflect.Type
			it      reflect.Type
			ip      *InputProfile
			ipState map[ProfileCategory]bool
		)
		if rh.refs[0].IsValid() {
			mr = rh.refs[0]
			mrt = mr.Type()
			it = mrt.In(1)
			ip = GetInputProfile(it)
			ipState = ip.State()
		}

		// stream request caller.
		var (
			smr      reflect.Value
			smrt     reflect.Type
			sit      reflect.Type
			sip      *InputProfile
			sipState map[ProfileCategory]bool
		)
		if len(rh.refs) == 2 && rh.refs[1].IsValid() {
			smr = rh.refs[1]
			smrt = smr.Type()
			sit = smrt.In(1)
			sip = GetInputProfile(sit)
			sipState = sip.State()
		}

		// construct virtual handler.
		var vh = func(c *gin.Context) {
			// auth request.
			var s = session.LoadSubject(c)
			if !s.Enforce(c, resource) {
				if s.IsAnonymous() {
					c.AbortWithStatus(http.StatusUnauthorized)
				} else {
					c.AbortWithStatus(http.StatusForbidden)
				}
				return
			}
			session.StoreSubjectCurrentOperation(c, s.Give(resource).If(c.Request.Method))

			// check request whether to stream.
			var withStream bool
			if rh.method == http.MethodGet && isUpgradeStreamRequest(c) {
				if !smr.IsValid() {
					c.AbortWithStatus(http.StatusBadRequest)
					return
				}
				withStream = true
			} else if !mr.IsValid() {
				c.AbortWithStatus(http.StatusBadRequest)
				return
			}

			// bind input.
			var (
				rmr      = mr
				rit      = it
				ripState = ipState
			)
			if withStream {
				rmr = smr
				rit = sit
				ripState = sipState
			}
			var ri reflect.Value
			if rit.Kind() == reflect.Pointer {
				ri = reflect.New(rit.Elem())
			} else {
				ri = reflect.New(rit)
			}
			if ripState[ProfileCategoryHeader] {
				if err := c.BindHeader(ri.Interface()); err != nil {
					return
				}
			}
			if ripState[ProfileCategoryUri] {
				if err := c.BindUri(ri.Interface()); err != nil {
					return
				}
			}
			if ripState[ProfileCategoryQuery] {
				if err := binding.MapFormWithTag(ri.Interface(), c.Request.URL.Query(), "query"); err != nil {
					return
				}
			}
			if c.Request.ContentLength != 0 {
				switch c.ContentType() {
				case binding.MIMEPOSTForm:
					if rh.method == http.MethodPost {
						if err := c.MustBindWith(ri.Interface(), binding.Form); err != nil {
							return
						}
					}
				case binding.MIMEMultipartPOSTForm:
					if rh.method == http.MethodPost {
						if err := c.MustBindWith(ri.Interface(), binding.FormMultipart); err != nil {
							return
						}
					}
				default:
					if rh.method != http.MethodGet {
						if err := c.BindJSON(ri.Interface()); err != nil {
							return
						}
					}
				}
			}
			if rv, ok := ri.Interface().(Validator); ok {
				if err := rv.Validate(); err != nil {
					_ = c.Error(err).
						SetType(gin.ErrorTypeBind)
					return
				}
			}
			if rv, ok := ri.Interface().(ValidatorWithInput); ok {
				hv, ok := h.(ValidatingInput)
				if ok {
					if err := rv.ValidateWith(c, hv.Validating()); err != nil {
						_ = c.Error(err).
							SetType(gin.ErrorTypeBind)
						return
					}
				}
			}
			if rit.Kind() != reflect.Pointer {
				ri = ri.Elem()
			}

			// process as stream request.
			if withStream {
				doUpgradeStreamRequest(c, rmr, ri)
				return
			}

			// process as normal request.
			var inputs = make([]reflect.Value, 0, 2)
			inputs = append(inputs, reflect.ValueOf(c))
			inputs = append(inputs, ri)
			var outputs = rmr.Call(inputs)

			// render response
			if c.Request.Context().Err() != nil ||
				c.Writer.Size() >= 0 ||
				len(c.Errors) != 0 {
				// already render inside the above processing
				return
			}
			var errInterface = outputs[len(outputs)-1].Interface()
			if errInterface != nil {
				var err = errInterface.(error)
				var ge = c.Error(err)
				if !isGinError(err) {
					_ = ge.SetType(gin.ErrorTypePublic)
				}
				return
			}
			var code = http.StatusOK
			switch len(outputs) {
			case 1:
				if rh.method == http.MethodPost {
					code = http.StatusNoContent
				}
				c.Writer.WriteHeader(code)
			case 2:
				if outputs[0].IsZero() {
					if rh.method == http.MethodPost {
						code = http.StatusNoContent
					}
					c.Writer.WriteHeader(code)
					return
				}
				if rh.method == http.MethodPost {
					code = http.StatusCreated
				}
				var obj = outputs[0].Interface()
				switch ot := obj.(type) {
				case rendCloser:
					if ot == nil {
						return
					}
					defer func() { _ = ot.Close() }()
					c.Render(code, ot)
				case render.Render:
					if ot == nil {
						return
					}
					c.Render(code, ot)
				default:
					c.JSON(code, obj) // TODO negotiate
				}
			case 3:
				var obj = GetResponseCollection(c, outputs[0].Interface(), int(outputs[1].Int()))
				c.JSON(code, obj) // TODO negotiate
			}
		}

		// register router.
		r.Handle(rh.method, rh.path, vh)

		// register schema.
		// TODO(thxCode) scheming Websocket.
		var (
			rip  = ip
			rmrt = mrt
		)
		if rip == nil {
			rip = sip
			rmrt = smrt
		}
		var rop = GetOutputProfile(rmrt)
		schemeRoute(resource, k+"."+rh.name, rh.method, rh.path, rip, rop)
	}

	if adv, ok := r.(AdviceAfterResourceRegistering); ok {
		var err = adv.AfterAdvice(adviceResource{Resource: h})
		if err != nil {
			return err
		}
	}

	return nil
}

type routeHandler struct {
	name   string
	method string
	path   string
	refs   []reflect.Value
}

const (
	createPrefix           = "Create"
	deletePrefix           = "Delete"
	updatePrefix           = "Update"
	getPrefix              = "Get"
	collectionCreatePrefix = "CollectionCreate"
	collectionDeletePrefix = "CollectionDelete"
	collectionUpdatePrefix = "CollectionUpdate"
	collectionGetPrefix    = "CollectionGet"

	routePrefix           = "Route"
	collectionRoutePrefix = "CollectionRoute"

	streamPrefix           = "Stream"
	collectionStreamPrefix = "CollectionStream"
)

func getRouteHandlers(h Resource, p string) []routeHandler {
	var logger = log.WithName("restful")

	var list []routeHandler
	var index = map[string]int{}

	var hr = reflect.ValueOf(h)
	var ht = hr.Type()

	var ms = sets.Set[string]{}
	for i := 0; i < ht.NumMethod(); i++ {
		ms.Insert(ht.Method(i).Name)
	}

	for _, mn := range sets.List[string](ms) {
		var mr = hr.MethodByName(mn)
		var mrt = mr.Type()

		var (
			rh            routeHandler
			forCollection bool
			forStream     bool
		)
		rh.name = mn

		// filter
		switch {
		default:
			continue
		case strings.HasPrefix(mn, streamPrefix):
			// attach to GET method.
			forStream = true
			rh.method = http.MethodGet
			rh.refs = []reflect.Value{{}, mr}
			if mn == streamPrefix {
				rh.path = path.Join("/", p, ":id")
			} else {
				rh.path = path.Join("/", p, ":id", strs.Dasherize(mn[len(streamPrefix):]))
			}
		case strings.HasPrefix(mn, collectionStreamPrefix):
			// attach to GET method.
			forStream = true
			forCollection = true
			rh.method = http.MethodGet
			rh.refs = []reflect.Value{{}, mr}
			if mn == collectionStreamPrefix {
				rh.path = path.Join("/", p)
			} else {
				rh.path = path.Join("/", p, "_", strs.Dasherize(mn[len(collectionStreamPrefix):]))
			}
		case strings.HasPrefix(mn, createPrefix):
			rh.method = http.MethodPost
			if mn == createPrefix {
				rh.path = path.Join("/", p)
			} else {
				rh.path = path.Join("/", p, ":id", strs.Dasherize(mn[len(createPrefix):]))
			}
		case strings.HasPrefix(mn, collectionCreatePrefix):
			forCollection = true
			rh.method = http.MethodPost
			if mn == collectionCreatePrefix {
				rh.path = path.Join("/", p, "_", "batch")
			} else {
				rh.path = path.Join("/", p, "_", strs.Dasherize(mn[len(collectionCreatePrefix):]))
			}
		case strings.HasPrefix(mn, deletePrefix):
			rh.method = http.MethodDelete
			if mn == deletePrefix {
				rh.path = path.Join("/", p, ":id")
			} else {
				rh.path = path.Join("/", p, ":id", strs.Dasherize(mn[len(deletePrefix):]))
			}
		case strings.HasPrefix(mn, collectionDeletePrefix):
			forCollection = true
			rh.method = http.MethodDelete
			if mn == collectionDeletePrefix {
				rh.path = path.Join("/", p)
			} else {
				rh.path = path.Join("/", p, "_", strs.Dasherize(mn[len(collectionDeletePrefix):]))
			}
		case strings.HasPrefix(mn, updatePrefix):
			rh.method = http.MethodPut
			if mn == updatePrefix {
				rh.path = path.Join("/", p, ":id")
			} else {
				rh.path = path.Join("/", p, ":id", strs.Dasherize(mn[len(updatePrefix):]))
			}
		case strings.HasPrefix(mn, collectionUpdatePrefix):
			forCollection = true
			rh.method = http.MethodPut
			if mn == collectionUpdatePrefix {
				rh.path = path.Join("/", p)
			} else {
				rh.path = path.Join("/", p, "_", strs.Dasherize(mn[len(collectionUpdatePrefix):]))
			}
		case strings.HasPrefix(mn, getPrefix):
			rh.method = http.MethodGet
			if mn == getPrefix {
				rh.path = path.Join("/", p, ":id")
			} else {
				rh.path = path.Join("/", p, ":id", strs.Dasherize(mn[len(getPrefix):]))
			}
		case strings.HasPrefix(mn, collectionGetPrefix):
			forCollection = true
			rh.method = http.MethodGet
			if mn == collectionGetPrefix {
				rh.path = path.Join("/", p)
			} else {
				rh.path = path.Join("/", p, "_", strs.Dasherize(mn[len(collectionGetPrefix):]))
			}
		}
		ms.Delete(mn)

		// validate
		var err = func() error {
			// validate input arguments.
			switch mrt.NumIn() {
			default:
				return fmt.Errorf("invalid input number of '%s': got %d but expected 2", mn, mrt.NumIn())
			case 2:
				if forStream {
					if !isTypeOfRequestStream(mrt.In(0)) {
						return fmt.Errorf("illegal first input type of '%s': expected runtime.RequestStream", mn)
					}
				} else {
					if !isTypeOfGinContext(mrt.In(0)) {
						return fmt.Errorf("illegal first input type of '%s': expected *gin.Context", mn)
					}
				}
				if !forCollection {
					switch mrt.In(1).Kind() {
					default:
						return fmt.Errorf("illegal last input type of '%s': expected struct or pointer type", mn)
					case reflect.Struct, reflect.Pointer:
					}
				} else {
					switch mrt.In(1).Kind() {
					default:
						return fmt.Errorf("illegal last input type of '%s': expected struct, pointer or slice type", mn)
					case reflect.Struct, reflect.Pointer, reflect.Slice:
					}
				}
			}
			// validate output arguments.
			switch mrt.NumOut() {
			default:
				return fmt.Errorf("invalid output number of '%s': got %d but expected not more than 3", mn, mrt.NumOut())
			case 3:
				if forStream || rh.method != http.MethodGet {
					return fmt.Errorf("invalid output number of '%s': got %d but exepcted not more than 2", mn, mrt.NumOut())
				}
				if !isTypeOfInt(mrt.Out(1)) {
					return fmt.Errorf("illegal second output type of '%s': expected int", mn)
				}
				if !isImplementationOfError(mrt.Out(2)) {
					return fmt.Errorf("illegal last output type of '%s': expected error", mn)
				}
			case 2:
				if forStream {
					return fmt.Errorf("invalid output number of '%s': got %d but exepcted not more than 1", mn, mrt.NumOut())
				}
				if !isImplementationOfError(mrt.Out(1)) {
					return fmt.Errorf("illegal last output type of '%s': expected error", mn)
				}
			case 1:
				if !forStream && rh.method == http.MethodGet {
					return fmt.Errorf("invalid output number of '%s': got %d but expected not more than 1", mn, mrt.NumOut())
				}
				if !isImplementationOfError(mrt.Out(0)) {
					return fmt.Errorf("illegal last output type of '%s': expected error", mn)
				}
			}
			return nil
		}()
		if err != nil {
			logger.Error(err)
			continue
		}

		rh.refs = []reflect.Value{mr}
		if rh.method == http.MethodGet {
			var idx, exist = index[rh.method+":"+rh.path]
			if exist {
				if forStream { // attach to GET handler.
					list[idx].refs = append(list[idx].refs, mr)
				} else {
					list[idx].refs[0] = mr
					list[idx].name = mn // rename to GET handler.
				}
				continue
			}
			if forStream { // attach to GET handler.
				rh.refs = []reflect.Value{{}, mr}
			}
		}
		index[rh.method+":"+rh.path] = len(list)
		list = append(list, rh)
	}

	for _, mn := range sets.List[string](ms) {
		var mr = hr.MethodByName(mn)
		var mrt = mr.Type()

		var (
			rh            routeHandler
			forCollection bool
		)
		rh.name = mn

		// filter
		switch {
		default:
			continue
		case strings.HasPrefix(mn, routePrefix):
			rh.path = path.Join("/", p, ":id") // part
		case strings.HasPrefix(mn, collectionRoutePrefix):
			forCollection = true
			rh.path = path.Join("/", p, "_") // part
		}
		ms.Delete(mn)

		// validate and complete
		var err = func() error {
			// validate input arguments.
			switch mrt.NumIn() {
			default:
				return fmt.Errorf("invalid input number of '%s': got %d but expected 2", mn, mrt.NumIn())
			case 2:
				if !isTypeOfGinContext(mrt.In(0)) {
					return fmt.Errorf("illegal first input type of '%s': expected *gin.Context", mn)
				}
				switch mrt.In(1).Kind() {
				default:
					return fmt.Errorf("illegal last input type of '%s': expected struct or pointer type", mn)
				case reflect.Struct, reflect.Pointer:
				}

				var rp = getProfileRouter(mrt.In(1))
				if rp == nil {
					return errors.New("illegal route profile: not found")
				}
				rh.method = rp.Method
				rh.path = path.Join(rh.path, rp.SubPath)
				if _, exist := index[rh.method+":"+rh.path]; exist {
					return fmt.Errorf("invalid subpath definition of '%s': conflict", mn)
				}
			}
			// validate output arguments.
			switch mrt.NumOut() {
			default:
				return fmt.Errorf("invalid output number of '%s': got %d but expected not more than 3", mn, mrt.NumOut())
			case 3:
				if !forCollection || rh.method != http.MethodGet {
					return fmt.Errorf("invalid output number of '%s': got %d but exepcted not more than 2", mn, mrt.NumOut())
				}
				if !isTypeOfInt(mrt.Out(1)) {
					return fmt.Errorf("illegal second output type of '%s': expected int", mn)
				}
				if !isImplementationOfError(mrt.Out(2)) {
					return fmt.Errorf("illegal last output type of '%s': expected error", mn)
				}
			case 2:
				if !isImplementationOfError(mrt.Out(1)) {
					return fmt.Errorf("illegal last output type of '%s': expected error", mn)
				}
			case 1:
				if rh.method == http.MethodGet {
					return fmt.Errorf("invalid output number of '%s': got %d but expected not more than 1", mn, mrt.NumOut())
				}
				if !isImplementationOfError(mrt.Out(0)) {
					return fmt.Errorf("illegal last output type of '%s': expected error", mn)
				}
			}
			return nil
		}()
		if err != nil {
			logger.Error(err)
			continue
		}

		rh.refs = []reflect.Value{mr}
		if rh.method == http.MethodGet {
			var idx, exist = index[rh.method+":"+rh.path]
			if exist {
				list[idx].refs[0] = mr
				list[idx].name = mn // rename to the GET handler.
				continue
			}
		}
		index[rh.method+":"+rh.path] = len(list)
		list = append(list, rh)
	}

	return list
}

type rendCloser interface {
	io.Closer
	render.Render
}

func isImplementationOfError(r reflect.Type) bool {
	var expected = reflect.TypeOf((*error)(nil)).Elem()
	switch r.Kind() {
	default:
		return false
	case reflect.Interface:
		return r.Implements(expected)
	}
}

func isTypeOfRequestStream(r reflect.Type) bool {
	var expected = reflect.TypeOf(RequestStream{}).String()
	switch r.Kind() {
	default:
		return false
	case reflect.Struct:
		var given = r.String()
		return given == expected
	}
}

func isTypeOfGinContext(r reflect.Type) bool {
	var expected = reflect.TypeOf(gin.Context{}).String()
	switch r.Kind() {
	default:
		return false
	case reflect.Pointer:
		var given = r.Elem().String()
		return given == expected
	}
}

func isTypeOfInt(r reflect.Type) bool {
	return r.Kind() == reflect.Int
}

func getResourceAndResourcePath(k string) (rk, rp string) {
	rk = strs.CamelizeDownFirst(strs.Pluralize(k))
	rp = strings.ToLower(strs.Pluralize(strs.Dasherize(rk)))
	return
}

type adviceResource struct {
	Resource
}

func (r adviceResource) ResourceAndResourcePath() (resource, resourcePath string) {
	return getResourceAndResourcePath(r.Kind())
}

func (r adviceResource) Unwrap() Resource {
	return r.Resource
}
