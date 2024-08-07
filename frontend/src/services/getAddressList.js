import dataArr from '../mock/dataArr.json';


const getAddressList = (filters) => {
    let list = dataArr.data.list;
    
    if (filters.accountNumber)
        list = list.filter(item => item.account.includes(filters.accountNumber));

    if (filters.is_active != undefined && filters.is_active != "all") {        
        list = list.filter(item => filters.is_active === item.is_active);        
    }

    if (!filters.is_organization != undefined && filters.is_organization != "all") {
        if (filters.is_organization === true) {
            list = list.filter(item => item.is_organization === filters.is_organization);
        } else {
            list = list.filter(item => item.is_organization === filters.is_organization || item.is_organization === undefined);
        }
    }

    if (!filters.is_residential != undefined && filters.is_residential != "all") {
        list = list.filter(item => item.is_residential == filters.is_residential);
    }

    if (filters.company != null) {
        list = list.filter(item => filters.company.includes(item.company.toString()));
    }
    
    return list.map(item => item.address);
}

export default getAddressList;