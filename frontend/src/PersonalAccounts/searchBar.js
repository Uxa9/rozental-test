import React from 'react';
import { Button, Input, Select, Space } from 'antd';

const { Search } = Input;


const SearchBar = (props) => {

    return (
        <Space.Compact>
            <Search 
                placeholder="Поиск по номеру счета..." 
                allowClear 
                onSearch={(e) => props.handleValue(e)}
            />
        </Space.Compact>
    )
}

export default SearchBar;