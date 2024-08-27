import styles from "./List.module.scss"

type Props = {
    elements: any[]
}

const List: React.FC<Props> = (props) => {
    
    return (
        <div className={styles["list-layout"]}>
            {props.elements.map((element, index) => {
                return (
                    <div key={index} className={styles["list-element"]}>
                        {element}
                    </div>
                )
            })}
        </div>
    )
}

export default List