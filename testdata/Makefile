.PHONY: all
all: $(addsuffix .json,\
	100 \
	500 \
	1000 \
	5000 \
	10000 \
	50000 \
	100000 \
	500000 \
)

%.json: student.yaml names.txt faculties.txt specialities.txt
	jg $< \
	  -a $* \
	  -f names=names.txt,faculties=faculties.txt,specialities=specialities.txt \
	  -o $@
