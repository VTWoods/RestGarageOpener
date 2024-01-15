#Golang Options
GOARCH=arm
GOOS=linux
GOARM=6

#File Options
BIN=garage_web
TMP=/tmp/${BIN}
OUT=bin/${BIN}
INSTALL_BIN=/usr/bin/${BIN}

all: ${OUT}

${OUT}: main.go
	GOARCH=${GOARCH} GOOS=${GOOS} GOARM=${GOARM} go build -o ${OUT} main.go

sync:
	rsync -avz ${OUT} ${TARGET}:${TMP}

run: sync
	ssh -t ${USER}@${TARGET} "${TMP} --address=0.0.0.0:8083" || true 

deploy: sync
	ssh -t ${USER}@${TARGET} "sudo cp ${TMP} ${INSTALL_BIN}; sudo service garage-web restart"
