package resolvers

import "cloud.google.com/go/firestore"

func (f *Resolver) Close(ticketId, reason string) error {
	_, err := f.firestore.Collection("reports").Doc(ticketId).Update(f.ctx, []firestore.Update{
		{
			Path:  "resolved",
			Value: true,
		},
		{
			Path:  "resolvedReason",
			Value: reason,
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
