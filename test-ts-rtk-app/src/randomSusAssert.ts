import * as images from "./components/images"

export const GetRandomSusById = (id: number) => {

    switch (id % 6) {
        case 0:
            return images.Black
        case 1:
            return images.White
        case 2:
            return images.Yellow
        case 3:
            return images.Red
        case 4:
            return images.Green
        case 5:
            return images.Blue
        
    }
}