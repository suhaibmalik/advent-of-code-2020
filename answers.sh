for d in day* ; do
    cd ${d}
    go run main.go > solutions.out
    cd ../
done