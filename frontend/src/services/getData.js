import dataArr from '../mock/dataArr.json';


const getData = (filters, pagination) => {
    let list = dataArr.data.list;
    console.log(filters);
    
    if (filters.accountNumber)
        list = list.filter(item => item.account.includes(filters.accountNumber));

    if (filters.is_active != undefined && filters.is_active != "all") {        
        list = list.filter(item => filters.is_active === item.is_active);        
    }

    if (filters.is_organization != undefined && filters.is_organization != "all") {
        if (filters.is_organization === true) {
            list = list.filter(item => item.is_organization === filters.is_organization);
        } else {
            list = list.filter(item => item.is_organization === filters.is_organization || item.is_organization === undefined);
        }
    }

    if (filters.is_residential != undefined && filters.is_residential != "all") {
        list = list.filter(item => item.is_residential == filters.is_residential);
    }

    if (filters.company != null) {
        list = list.filter(item => filters.company.includes(item.company.toString()));
    }

    if (filters.apartment != null) {
        list = list.filter(item => item.apartment.includes(filters.apartment));
    }

    if (filters.address != null) {
        list = list.filter(item => item.address.includes(filters.address));
    }
    
    console.log(list);
    const total = list.length

    if (pagination) list = list.slice((pagination.page-1) * pagination.perPage, (pagination.page-1) * pagination.perPage + pagination.perPage);
    
    return {list, total}
}

export default getData;