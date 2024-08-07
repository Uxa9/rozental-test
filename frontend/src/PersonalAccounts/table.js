import React, { useEffect, useState } from 'react';
import { Empty, Table } from 'antd';
import ModalAccountView from './modalAccountView';

const tableColumns = [
    {
        title: 'Лицевой счет',
        dataIndex: 'account',
        fixed: 'left',
    },
    {
        title: 'Адрес',
        dataIndex: 'address'
    },
    {
        title: 'Помещение',
        dataIndex: 'apartment'
    },
    {
        title: "Назначение помещения",
        dataIndex: 'is_residential',
        render: (type) => type ? "Жилое" : "Не жилое"
    },
    {
        title: 'ФИО',
        dataIndex: 'customer',
        key: 'name',
        render: (customer) => customer?.fio
    },
    {
        title: 'Телефон',
        dataIndex: 'customer',
        key: 'phone',
        render: (customer) => <a href={`tel:${customer.phone}`}>{customer.phone}</a>
    }
]


const AccountsTable = (props) => {

    const [data, setData] = useState({list: [], total: 0});
    const [tablePagination, setTablePagination] = useState([1, 10]);
    const [openModal, setOpenModal] = useState(false);
    const [modalData, setModalData] = useState({});

    useEffect(() => {
        setData(props.data || []);
    }, [props.data]);

    useEffect(() => {
        props.handlePaginationChanges(tablePagination);
    }, [tablePagination]);

    return (
        <>
            <ModalAccountView 
                open={openModal} 
                setOpenModal={setOpenModal}
                data={modalData}
            />
            <Table 
                scroll={{
                    x: 800,
                }} 
                columns={tableColumns} 
                dataSource={data.list}
                pagination={{
                    total: data.total,
                    current: tablePagination[0],
                    pageSize: tablePagination[1],
                    onChange: (page, pageSize) => {
                        setTablePagination([page, pageSize])
                    }
                }}
                rowClassName={"row-cursor"}
                onRow={(record, rowIndex) => {                    
                    return {
                        onClick: () => {
                            setOpenModal(true);
                            setModalData(record);
                        }, 
                    };
                }}
            />
        </>
    );
}

export default AccountsTable;