package resolvers

import "cloud.google.com/go/firestore"

func (f *Resolver) Reopen(ticketId string) error {
	_, err := f.firestore.Collection("reports").Doc(ticketId).Update(f.ctx, []firestore.Update{
		{
			Path:  "resolved",
			Value: false,
		},
		{
			Path:  "resolvedReason",
			Value: "",
		},
		{
			Path:  "updatedAt",
			Value: firestore.ServerTimestamp,
		},
	})
	if err != nil {
		return err
	}
	return nil
}
