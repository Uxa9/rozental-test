import { useEffect, useState } from "react";
import { Product } from "../../common/productStore/product";
import Button from "../button/Button";
import styles from "./ProductForm.module.scss"

type Props = {
    product?: Product,
    resultHandler: (data: Product) => void
}

const ProductForm: React.FC<Props> = (props) => {

    const [data, setData] = useState<Product>({
        id: 0,
        name: "",
        price: 0
    })

    useEffect(() => {
        if (props.product !== undefined) {
            setData({
                id: props.product.id,
                name: props.product.name,
                price: props.product.price
            })
        }
    }, [props.product])

    const editName = (event: React.ChangeEvent<HTMLInputElement>) => {
        setData({
            id: data.id,
            name: event.target.value,
            price: data.price
        })
    }

    const editPrice = (event: React.ChangeEvent<HTMLInputElement>) => {
        setData({
            id: data.id,
            name: data.name,
            price: Number(event.target.value)
        })
    }

    const handleFormSubmit = (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault()        
        props.resultHandler(data)        
    }

    return (
        <form onSubmit={handleFormSubmit} className={styles["form-wrapper"]}>
            <label>Название продукта:</label>
            <input type="text" name="name" value={data.name || undefined} onChange={editName} placeholder="Введите название" />
            <label>Цена продукта:</label>
            <input type="number" name="price" value={data.price || undefined} onChange={editPrice} placeholder="Введите цену" />
            <Button text={!props.product ? "Добавить продукт" : "Изменить"} type="submit" />
        </form>
    )
}

export default ProductForm;