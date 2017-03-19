
service buf {
    void WriteData(1: string filename, 2: string data)
    string ReadData(1: string filename)
}
