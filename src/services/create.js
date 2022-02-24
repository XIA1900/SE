import { request } from 'umi';

export async function createGroup(params) {   
    return request('/api/createGroup', {
      data: params,
      method: 'POST',
    });
}

export async function createPost(params) {  
  return request('/api/createPost', {
    data: params,
    method: 'POST',
  });
}