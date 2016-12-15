// Gowurfl is a wrapper around libwurfl from scientiamobile.
// To use this package you need to have both the headers and
// the library in a place that cgo can find them, e.g /usr/lib/
// and /usr/include.
package gowurfl

// #cgo LDFLAGS: -lwurfl
/*
#include <wurfl/wurfl.h>
#include <stdlib.h>
*/
import "C"
import "unsafe"

import (
	"errors"
	"strconv"
)

func New() (*WURFL, error) {
	h := C.wurfl_handle(C.wurfl_create())
	w := &WURFL{handle: h}

	if err := w.CheckError(); err != nil {
		return nil, err
	}

	return w, nil
}

type WURFL struct {
	handle C.wurfl_handle
	path   string
}

type WURFLError error

// s/_\(\w\)\([^_]\+\)/\1\L\2/g
var (
	ErrorInvalidHandle                     = errors.New("handle passed to the function is invalid")
	ErrorAlreadyLoad                       = errors.New("wurflload has already been invoked on the specific wurflhandle")
	ErrorFileNotFound                      = errors.New("file not found during wurflload")
	ErrorUnexpectedEndOfFile               = errors.New("unexpected end of file or parsing error during wurflload")
	ErrorInputOutputFailure                = errors.New("error reading stream during wurflload")
	ErrorDeviceNotFound                    = errors.New("specified device is missing")
	ErrorCapabilityNotFound                = errors.New("specified capability is missing")
	ErrorInvalidCapabilityValue            = errors.New("invalid capability value")
	ErrorVirtualCapabilityNotFound         = errors.New("specified virtual capability is missing")
	ErrorCantLoadCapabilityNotFound        = errors.New("specified capability is missing")
	ErrorCantLoadVirtualCapabilityNotFound = errors.New("specified virtual capability is missing")
	ErrorEmptyID                           = errors.New("missing id in searching device")
	ErrorCapabilityGroupNotFound           = errors.New("specified capability is missing in its group")
	ErrorCapabilityGroupMismatch           = errors.New("specified capability mismatch in its group")
	ErrorDeviceAlreadyDefined              = errors.New("specified device is already defined")
	ErrorUseragentAlreadyDefined           = errors.New("specified user agent is already defined")
	ErrorDeviceHierarchyCircularReference  = errors.New("circular reference in device hierarchy ")
	ErrorUnknown                           = errors.New("unknown error")
	ErrorInvalidUseragentPriority          = errors.New("specified override sideloaded browser user agent configuration not valid")
	ErrorInvalidParameter                  = errors.New("invalid parameter")
	ErrorInvalidCacheSize                  = errors.New("specified an invalid cache size, 0 or a negative value.")
	ErrorXMLConsistency                    = errors.New("wurfl.xml is out of date - some needed deviceid is missing")
)

func goError(e C.wurfl_error) error {
	switch e {
	default:
		return ErrorUnknown
	case C.WURFL_ERROR_DEVICE_NOT_FOUND:
		return ErrorDeviceNotFound
	case C.WURFL_ERROR_CAPABILITY_NOT_FOUND:
		return ErrorCapabilityNotFound
	case C.WURFL_ERROR_INVALID_HANDLE:
		return ErrorInvalidHandle
	case C.WURFL_ERROR_ALREADY_LOAD:
		return ErrorAlreadyLoad
	case C.WURFL_ERROR_FILE_NOT_FOUND:
		return ErrorFileNotFound
	case C.WURFL_ERROR_UNEXPECTED_END_OF_FILE:
		return ErrorUnexpectedEndOfFile
	case C.WURFL_ERROR_INPUT_OUTPUT_FAILURE:
		return ErrorInputOutputFailure
	case C.WURFL_ERROR_INVALID_CAPABILITY_VALUE:
		return ErrorInvalidCapabilityValue
	case C.WURFL_ERROR_VIRTUAL_CAPABILITY_NOT_FOUND:
		return ErrorVirtualCapabilityNotFound
	case C.WURFL_ERROR_CANT_LOAD_CAPABILITY_NOT_FOUND:
		return ErrorCantLoadCapabilityNotFound
	case C.WURFL_ERROR_CANT_LOAD_VIRTUAL_CAPABILITY_NOT_FOUND:
		return ErrorCantLoadVirtualCapabilityNotFound
	case C.WURFL_ERROR_EMPTY_ID:
		return ErrorEmptyID
	case C.WURFL_ERROR_CAPABILITY_GROUP_NOT_FOUND:
		return ErrorCapabilityGroupNotFound
	case C.WURFL_ERROR_CAPABILITY_GROUP_MISMATCH:
		return ErrorCapabilityGroupMismatch
	case C.WURFL_ERROR_DEVICE_ALREADY_DEFINED:
		return ErrorDeviceAlreadyDefined
	case C.WURFL_ERROR_USERAGENT_ALREADY_DEFINED:
		return ErrorUseragentAlreadyDefined
	case C.WURFL_ERROR_DEVICE_HIERARCHY_CIRCULAR_REFERENCE:
		return ErrorDeviceHierarchyCircularReference
	case C.WURFL_ERROR_UNKNOWN:
		return ErrorUnknown
	case C.WURFL_ERROR_INVALID_USERAGENT_PRIORITY:
		return ErrorInvalidUseragentPriority
	case C.WURFL_ERROR_INVALID_PARAMETER:
		return ErrorInvalidParameter
	case C.WURFL_ERROR_INVALID_CACHE_SIZE:
		return ErrorInvalidCacheSize
	case C.WURFL_ERROR_XML_CONSISTENCY:
		return ErrorXMLConsistency
	}
}

func (w *WURFL) ErrorCount() int {
	return int(C.wurfl_has_error_message(w.handle))
}

func (w *WURFL) ClearErrors() {
	C.wurfl_clear_error_message(w.handle)
}

// LastError checks if libwurfl has any error messages queued
// up and retrieves the last one if that is the case.
func (w *WURFL) LastError() error {
	err := C.wurfl_get_error_message(w.handle)
	if err == nil {
		return errors.New("failed to get last error")
	}

	es := C.GoString(err)
	if es == "" {
		return nil
	}

	return errors.New(es)
}

// CheckError is a convenience method that checks the error count,
// retrieves it if > 0, clears them and then returns it.
func (w *WURFL) CheckError() error {
	if w.ErrorCount() > 0 {
		err := w.LastError()
		w.ClearErrors()
		return err
	}

	return nil
}

type EngineTarget int

const (
	EngineTargetHighAccuracy    EngineTarget = C.WURFL_ENGINE_TARGET_HIGH_ACCURACY
	EngineTargetHighPerformance EngineTarget = C.WURFL_ENGINE_TARGET_HIGH_PERFORMANCE
	EngineTargetInvalid         EngineTarget = C.WURFL_ENGINE_TARGET_INVALID
)

// GetEngineTarget retrieves the current engine target. The default target is
// EngineTargetHighPerformance.
func (w *WURFL) GetEngineTarget() EngineTarget {
	et := C.wurfl_get_engine_target(w.handle)

	return EngineTarget(et)
}

func (w *WURFL) SetEngineTarget(et EngineTarget) error {
	err := C.wurfl_error(C.wurfl_set_engine_target(w.handle, C.wurfl_engine_target(et)))

	if err != C.WURFL_OK {
		return goError(err)
	}

	return nil
}

type CacheProvider int

const (
	CacheProviderNone      CacheProvider = C.WURFL_CACHE_PROVIDER_NONE
	CacheProviderLRU       CacheProvider = C.WURFL_CACHE_PROVIDER_LRU
	CacheProviderDoubleLRU CacheProvider = C.WURFL_CACHE_PROVIDER_DOUBLE_LRU
)

const (
	defaultCacheProviderSize = "10000, 3000"
)

// SetCacheProvider sets the caching strategy to use.
// From the documentation:
//
// 		The default is Double LRU, which is a two-cache strategy (one going from
// 		User-Agent to Device-Id, the other from Device-Id to Device). The default
// 		parameters are 10,000, 3,000 (maximum 10,000 elements for the User-Agent
// 		to device-id cache and maximum 3,000 elements for the device-id to device
// 		cahse) and the values are in elements. The LRU cache comes with User-Agent
// 		to Device mapping only, and the NULL parameter will disable the cache mode.
//
// For the CacheProviderNone all size parameters are ignored, CacheProviderLRU
// only uses the first parameter and CacheProviderDoubleLRU uses two.
// CacheProviderDoubleLRU uses the default if there are not enough size parameters.
func (w *WURFL) SetCacheProvider(c CacheProvider, sizes ...int) error {
	var cfg *C.char
	defer C.free(unsafe.Pointer(cfg))

	switch c {
	default:
		return errors.New("invalid cache provider")
	case CacheProviderNone:
		cfg = nil
	case CacheProviderLRU:
		if len(sizes) >= 1 {
			cfg = C.CString(strconv.Itoa(sizes[0]))
		} else {
			cfg = C.CString(defaultCacheProviderSize)
		}
	case CacheProviderDoubleLRU:
		if len(sizes) >= 2 {
			s := strconv.Itoa(sizes[0]) + ", " + strconv.Itoa(sizes[1])
			cfg = C.CString(s)
		} else {
			cfg = C.CString(defaultCacheProviderSize)
		}
	}

	err := C.wurfl_set_cache_provider(w.handle, C.wurfl_cache_provider(c), cfg)
	if err != C.WURFL_OK {
		return goError(err)
	}

	return nil
}

func (w *WURFL) SetRoot(p string) error {
	ps := C.CString(p)
	defer C.free(unsafe.Pointer(ps))

	err := C.wurfl_set_root(w.handle, ps)

	if err != C.WURFL_OK {
		return goError(err)
	}

	w.path = p
	return nil
}

func (w *WURFL) GetInfo() (string, error) {
	r := C.wurfl_get_wurfl_info(w.handle)

	if r == nil {
		return "", errors.New("called GetInfo() before loading root file")
	}
	s := C.GoString(r)

	return s, nil
}

// Load performs WURFL root file and patch loading.
// It must be called after specifying root filename by calling SetRoot to set
// the root file name, and optionally AddPatch and/or AddRequestedCapability
// to respectively specify patches and requested capabilities (if no capability
// is requested, all capabilities from WURFL root file and patches are loaded).
func (w *WURFL) Load() error {
	err := C.wurfl_error(C.wurfl_load(w.handle))

	if err != C.WURFL_OK {
		return goError(err)
	}

	return nil
}

func (w *WURFL) Close() {
	C.wurfl_destroy(w.handle)
}

type Capabilities map[string]string

var (
	MandatoryCapabilities = []string{
		"device_os",
		"device_os_version",
		"is_tablet",
		"is_wireless_device",
		"pointing_method",
		"preferred_markup",
		"resolution_height",
		"resolution_width",
		"ux_full_desktop",
		"xhtml_support_level",
		"is_smarttv",
		"can_assign_phone_number",
		"brand_name",
		"model_name",
		"marketing_name",
		"mobile_browser_version",
	}
)

// AddRequestedCapability adds a capability to the "Requested Capabilities" collection.
// If this function is never called, the Load() function will automatically
// load all the features in the WURFL database.
// If one or more capabilities are added to the "Requested Capabilities"
// collection, only the specified capabilities (if available) will be loaded
// from the database, resulting in a reduced memory footprint.
func (w *WURFL) AddRequestedCapability(cap string) error {
	cc := C.CString(cap)
	defer C.free(unsafe.Pointer(cc))

	err := C.wurfl_add_requested_capability(w.handle, cc)
	if err != C.WURFL_OK {
		return goError(err)
	}

	return nil
}

// AddRequestedCapabilities is a convenience method to add a list of capabilities
// all at once.
func (w *WURFL) AddRequestedCapabilities(caps []string) error {
	for _, cap := range caps {
		if err := w.AddRequestedCapability(cap); err != nil {
			return err
		}
	}

	return nil
}

func (w *WURFL) HasCapability(cap string) bool {
	cc := C.CString(cap)
	defer C.free(unsafe.Pointer(cc))

	ret := C.wurfl_has_capability(w.handle, cc)

	return int(ret) > 0
}

func (w *WURFL) GetMandatoryCapabilities() ([]string, error) {
	caps := []string{}

	enum := C.wurfl_get_mandatory_capability_enumerator(w.handle)
	defer C.wurfl_capability_enumerator_destroy(enum)

	for C.wurfl_capability_enumerator_is_valid(enum) == 1 {
		name := C.wurfl_capability_enumerator_get_name(enum)
		if name == nil {
			return caps, errors.New("failed to get name for capability enumerator")
		}

		caps = append(caps, C.GoString(name))
		C.wurfl_capability_enumerator_move_next(enum)
	}

	return caps, nil
}

func (w *WURFL) GetCapabilities() ([]string, error) {
	caps := []string{}

	enum := C.wurfl_get_capability_enumerator(w.handle)
	defer C.wurfl_capability_enumerator_destroy(enum)

	for C.wurfl_capability_enumerator_is_valid(enum) == 1 {
		name := C.wurfl_capability_enumerator_get_name(enum)
		if name == nil {
			return caps, errors.New("failed to get name for capability enumerator")
		}

		caps = append(caps, C.GoString(name))
		C.wurfl_capability_enumerator_move_next(enum)
	}

	return caps, nil
}

type Device struct {
	handle C.wurfl_device_handle
}

func (w *WURFL) LookupUserAgent(ua string) (*Device, error) {
	cua := C.CString(ua)
	defer C.free(unsafe.Pointer(cua))
	h := C.wurfl_lookup_useragent(w.handle, cua)

	if h == nil {
		return nil, errors.New("failed to look up user agent")
	}

	return &Device{handle: h}, nil
}

func (d *Device) GetID() (string, error) {
	id := C.wurfl_device_get_id(d.handle)

	if id == nil {
		return "", errors.New("failed to query for device id")
	}

	return C.GoString(id), nil
}

func (d *Device) HasVirtualCapability(cap string) (bool, error) {
	cc := C.CString(cap)
	defer C.free(unsafe.Pointer(cc))
	c := int(C.wurfl_device_has_virtual_capability(d.handle, cc))

	if c == -1 {
		return false, errors.New("failed to lookup virtual capability")
	}

	return c == 1, nil
}

//
func (d *Device) GetVirtualCapability(cap string) (string, error) {
	cc := C.CString(cap)
	defer C.free(unsafe.Pointer(cc))
	c := C.wurfl_device_get_virtual_capability(d.handle, cc)

	if c == nil {
		return "", errors.New("failed to lookup virtual capability")
	}

	return C.GoString(c), nil
}

func (d *Device) GetVirtualCapabilities() (Capabilities, error) {
	caps := make(Capabilities)

	enum := C.wurfl_device_get_virtual_capability_enumerator(d.handle)
	defer C.wurfl_device_capability_enumerator_destroy(enum)

	for C.wurfl_device_capability_enumerator_is_valid(enum) == 1 {
		name := C.wurfl_device_capability_enumerator_get_name(enum)
		value := C.wurfl_device_capability_enumerator_get_value(enum)
		caps[C.GoString(name)] = C.GoString(value)
		C.wurfl_device_capability_enumerator_move_next(enum)
	}

	return caps, nil
}

func (d *Device) GetCapabilitiy(name string) (string, error) {
	cc := C.CString(name)
	defer C.free(unsafe.Pointer(cc))

	c := C.wurfl_device_get_capability(d.handle, cc)

	if c == nil {
		return "", errors.New("failed to lookup capability")
	}

	return C.GoString(c), nil
}

func (d *Device) Close() {
	C.wurfl_device_destroy(d.handle)
}
