	import hashlib
	seq = 1
	m = hashlib.md5()

	while not m.hexdigest().startswith('00000'):
		m = hashlib.md5()
		seq+=1
		m.update('iwrupvqb%d'%seq)

	print  "this is with five 0: ",seq

	seq = 0

	while not m.hexdigest().startswith('000000'):
		m = hashlib.md5()
		seq+=1
		m.update('iwrupvqb%d'%seq)

	print  "this is with six 0: ",seq