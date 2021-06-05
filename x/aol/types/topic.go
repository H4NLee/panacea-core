package types

func (t Topic) NextRecordOffset() uint64 {
	return t.TotalRecords
}

func (t Topic) IncreaseTotalRecords() Topic {
	return Topic{
		TotalRecords: t.TotalRecords + 1,
		TotalWriters: t.TotalWriters,
	}
}

func (t Topic) IncreaseTotalWriters() Topic {
	return Topic{
		TotalRecords: t.TotalRecords,
		TotalWriters: t.TotalWriters + 1,
	}
}

func (t Topic) DecreaseTotalWriters() Topic {
	return Topic{
		TotalRecords: t.TotalRecords,
		TotalWriters: t.TotalWriters - 1,
	}
}
