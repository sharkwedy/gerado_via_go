package model

type Product struct {
    ID   string `json:"id"`
    Name string `json:"name"`
    Data struct {
        Color         string  `json:"color"`
        Capacity      string  `json:"capacity"`
        Price         float64 `json:"price"`
        Generation    string  `json:"generation"`
        Year          int     `json:"year"`
        CpuModel      string  `json:"cpuModel"`
        HardDiskSize  string  `json:"hardDiskSize"`
        StrapColour   string  `json:"strapColour"`
        CaseSize      string  `json:"caseSize"`
        Description   string  `json:"description"`
        ScreenSize    float64 `json:"screenSize"`
    } `json:"data"`
}