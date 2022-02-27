function getNavItems(data, tofilter) {
    const n = Object.keys(data)
        .filter(key => tofilter.includes(key))
        .reduce((obj, key) => {
            obj[key] = data[key];
            return obj;
        }, {});
    return Object.values(n);
};

export { getNavItems }