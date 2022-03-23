import { request } from 'umi';

export async function getBasicInfo(params) {
    return request('/api/getBasicInfo', {
        params,
    });
}

export async function getAnalysis(params) {
    return request('/api/getAnalysis', {
        params,
    });
}

export async function getMember(params) {
    return request('/api/getMember', {
        params,
    });
}


export async function getNotification(params) {
    return request('/api/getNotification', {
        params,
    });
}

export async function updateGroupInfo(params) {
    return request('/api/updateGroupInfo', {
        data: params,
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
    });
}

export async function deleteGroup(params) {
    return request('/api/deleteGroup', {
        data: params,
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
    });
}

export async function deleteMember(params) {
    return request('/api/deleteMember', {
        data: params,
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
    });
}
  
