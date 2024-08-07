import React, { Children, useEffect, useState } from 'react';
import { Modal, Radio, Descriptions, Select, Input, Flex } from 'antd';
import dictionary from '../mock/dictionary.json';
import getAddressList from '../services/getAddressList';

const ModalFilterView = (props) => {
    const [isModalOpen, setIsModalOpen] = useState(false);
    const [filters, setFilters] = useState({
        is_active: 'all',
        is_organization: 'all',
        is_residential: 'all',
        company: null,
        address: null,
        apartment: null
    });
    const [addressList, setAddressList] = useState([getAddressList(filters)]);

    const showModal = () => {
        setIsModalOpen(true);
    };

    const handleOk = () => {
        setIsModalOpen(false);
        props.applyFilters(filters);
        props.closeModal(false);
    };

    const handleCancel = () => {
        setIsModalOpen(false);
        props.closeModal(false);
    };

    useEffect(() => {
        setIsModalOpen(props.open);
    }, [props.open]);

    useEffect(() => {
        setAddressList(getAddressList(filters))
    }, [filters]);

    const activeOptions = [
        { label: 'Все', value: 'all' },
        { label: 'Активные', value: true },
        { label: 'Закрытые', value: false },
    ];

    const ownerOptions = [
        { label: 'Все', value: 'all' },
        { label: 'Юр. лицо', value: true },
        { label: 'Физ. лицо', value: false },
    ];

    const buildingOptions = [
        { label: 'Все', value: 'all' },
        { label: 'Жилое', value: true },
        { label: 'Не жилое', value: false },
    ];

    const items = [
        {
            key: '1',
            label: 'Активность',
            children:
                <Radio.Group
                    options={activeOptions}
                    onChange={(e) =>
                        setFilters({ ...filters, is_active: e.target.value })
                    }
                    value={filters.is_active}
                    optionType="button"
                    buttonStyle="solid"
                    name='Aboba'
                />
        },
        {
            key: '2',
            label: 'Тип собственника',
            children:
                <Radio.Group
                    options={ownerOptions}
                    onChange={(e) =>
                        setFilters({ ...filters, is_organization: e.target.value })
                    }
                    value={filters.is_organization}
                    optionType="button"
                    buttonStyle="solid"
                />
        },
        {
            key: '3',
            label: 'Тип помещения',
            children:
                <Radio.Group
                    options={buildingOptions}
                    onChange={(e) =>
                        setFilters({ ...filters, is_residential: e.target.value })
                    }
                    value={filters.is_residential}
                    optionType="button"
                    buttonStyle="solid"
                />
        },
        {
            key: '4',
            label: 'Компании',
            children:
                <Flex>
                    <Select
                        showSearch
                        allowClear
                        mode='multiple'
                        style={{minWidth: "250px"}}
                        placeholder="Выберите компанию"
                        optionFilterProp="label"
                        onChange={(value) => setFilters({ ...filters, company: value })}
                        options={Object.keys(dictionary).map((key) => {
                            return {
                                value: key,
                                label: dictionary[key]
                            }
                        })}
                    />
                </Flex>
        },
        {
            key: '5',
            label: 'Дома',
            children:
                <Select
                    showSearch
                    allowClear
                    style={{minWidth: "250px"}}
                    placeholder="Выберите дом"
                    optionFilterProp="label"
                    onChange={(value) => setFilters({ ...filters, address: value })}
                    options={addressList.map((item, index) => {
                        return {
                            value: item,
                            key: index,
                            label: item
                        }
                    })}
                />
        },
        {
            key: '6',
            label: 'Помещение',
            children:
                <Input 
                    placeholder='Введите номер помещения'
                    onChange={(e) => setFilters({ ...filters, apartment: e.target.value })}
                />
        },
    ]

    return (
        <Modal
            title={`Фильтры`}
            open={isModalOpen}
            onOk={handleOk}
            onCancel={handleCancel}
            footer={(_, { OkBtn, CancelBtn }) => (
                <>
                    <OkBtn>
                        Применить
                    </OkBtn>
                </>
            )}
        >
            <Descriptions items={items} layout='vertical' column={1} />
        </Modal>
    )
}

export default ModalFilterView;