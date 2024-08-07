import React, { Children, useEffect, useState } from 'react';
import { Modal, Descriptions } from 'antd';
import dictionary from '../mock/dictionary.json';

const ModalAccountView = (props) => {
    const [isModalOpen, setIsModalOpen] = useState(false);
    const [data, setData] = useState([]);

    const showModal = () => {
        setIsModalOpen(true);
    };

    const handleOk = () => {
        setIsModalOpen(false);
        props.setOpenModal(false);
    };

    const handleCancel = () => {
        setIsModalOpen(false);
        props.setOpenModal(false);
    };

    useEffect(() => {
        setIsModalOpen(props.open);
    }, [props.open]);

    useEffect(() => {
        props.data && setData([
            {
                key: 'name',
                label: props.data?.is_organization ? "Название организации" : "ФИО",
                children: props.data?.customer?.fio
            },
            {
                key: 'address',
                label: 'Дом',
                children: props.data?.address
            },
            {
                key: 'houseNumber',
                label: 'Помещение',
                children: props.data?.apartment
            },
            {
                key: 'company',
                label: "Компания",
                children: dictionary[Object.keys(dictionary).find(key => key == props.data?.company)]
            },
            {
                key: 'phone',
                label: 'Телефон',
                children: props.data?.customer?.phone
            },
            {
                key: 'email',
                label: 'Email',
                children: props.data?.customer?.email
            }
        ])
    }, [props.data]);

    return (
        <Modal
            title={`Лицевой счет: ${props?.data.account}`}
            open={isModalOpen}
            onOk={handleOk}
            onCancel={handleCancel}
            footer={(_, { OkBtn, CancelBtn }) => (
                <>
                    <OkBtn />
                </>
            )}
        >
            <Descriptions title="Основные данные" items={data} layout='vertical' />
        </Modal>
    )
}

export default ModalAccountView;