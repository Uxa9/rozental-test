import React, { useEffect, useState } from 'react';
import AccountsTable from './table';
import getData from '../services/getData';
import SearchBar from './searchBar';
import ModalFilterView from './modalFilterView';
import { Button, Divider, Space } from 'antd';
import { FilterOutlined } from '@ant-design/icons';

const PersonalAccounts = () => {

    const [data, setData] = useState({list: [], total: 0});
    const [tablePagination, setTablePagination] = useState([1, 10]);
    const [accountNumber, setAccountNumber] = useState(null);
    const [filterModalOpen, setFilterModalOpen] = useState(false);
    const [filters, setFilters] = useState({})

    const handlePaginationChanges = ([page, pageSize]) => setTablePagination([page, pageSize]);
    const handleSearchBar = (number) => setAccountNumber(number);
    const applyFilters = () => {}

    useEffect(() => {        
        setData(getData({accountNumber, ...filters}, {page: tablePagination[0], perPage: tablePagination[1]}));
    }, [tablePagination, accountNumber, filters]);


    return (
        <>
            <ModalFilterView 
                open={filterModalOpen}
                applyFilters={setFilters}
                closeModal={setFilterModalOpen}
            />
            <Space size={'middle'}>
                <SearchBar 
                    handleValue={handleSearchBar}
                />
                <Button 
                    icon={<FilterOutlined />}
                    onClick={() => setFilterModalOpen(true)}
                >
                    Фильтры
                </Button>
            </Space>
            <Divider />
            <AccountsTable 
                data={data}
                handlePaginationChanges={handlePaginationChanges}
            />
        </>
    );
}

export default PersonalAccounts;