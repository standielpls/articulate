deploy:
	gcloud functions deploy testMe --entry-point Handler --runtime go111 --trigger-http --allow-unauthenticated
	gcloud functions deploy createNote --entry-point Handler --runtime go111 --trigger-http --allow-unauthenticated
	gcloud functions deploy listNotesByUserID --entry-point Handler --runtime go111 --trigger-http --allow-unauthenticated

