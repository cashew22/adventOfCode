class gift:
	def __init__(self, size):
		self.l = size[0]
		self.w = size[1]
		self.h = size[2]
		self.lw = self.l * self.w
		self.wh = self.w * self.h
		self.hl = self.h * self.l
		self.paper = min(self.lw, self.wh, self.hl) + 2*(self.lw + self.wh + self.hl)
		self.bow = self.l * self.w * self.h
		size.sort()
		self.ribbon = 2 * size[0] + 2 * size[1]

f = open('input.txt', 'r')
paper = 0
ribbon = 0

for l in f:
	size = [int(x) for x in l.split('x')]
	mygift = gift(size)
	paper += mygift.paper
	ribbon += mygift.bow + mygift.ribbon
print "You should buy that amount of paper: ", paper
print "and you should buy that amount of ribbon: ", ribbon