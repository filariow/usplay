package datext

import (
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// ExtractDatesStr extracts printable representation of dates from intervals
func ExtractDatesStr(period interface {
	GetFrom() *timestamppb.Timestamp
	GetTo() *timestamppb.Timestamp
}) (from, to string) {
	_from, err := ptypes.Timestamp(period.GetFrom())
	if err != nil {
		from = period.GetFrom().String()
	} else {
		from = _from.String()
	}

	_to, err := ptypes.Timestamp(period.GetTo())
	if err != nil {
		to = period.GetTo().String()
	} else {
		to = _to.String()
	}
	return
}
