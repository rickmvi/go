package collections

import (
	"github.com/rickmvi/go/pkg/collections/map"
	"github.com/rickmvi/go/pkg/util/function"
	"github.com/rickmvi/go/pkg/util/function/streams"
	_type "github.com/rickmvi/go/pkg/util/type"
)

func ToMap[T _type.Any, K comparable, V any](
	s *streams.Stream[T],
	keyMapper function.Function[T, K],
	valueMapper function.Function[T, V],
	mergeFunction function.BiFunction[V, V, V],
) *_map.Map[K, V] {

	resultMap := _map.New[K, V]()

	s.ForEach(func(element T) {
		key := keyMapper(element)
		value := valueMapper(element)

		if !resultMap.ContainsKey(key) {
			resultMap.Put(key, value)
			return
		}

		oldValue, _ := resultMap.Get(key)

		if mergeFunction == nil {
			resultMap.Put(key, value)
			return
		}

		mergedValue := mergeFunction(oldValue, value)
		resultMap.Put(key, mergedValue)
	})

	return resultMap
}
