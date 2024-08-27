import styles from "./Table.module.scss"

type Props = {
    elements: any[]
}

const Table: React.FC<Props> = (props) => {
    
    return (
        <div className={styles["table-layout"]}>
            {props.elements.map((element, index) => {
                return (
                    <div key={String(index)} className={styles["table-element"]}>
                        {element}
                    </div>
                )
            })}
        </div>
    )
}

export default Table