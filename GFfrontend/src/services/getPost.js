import { request } from 'umi';

export async function getPost(params) {
  //ID: article id
  //username
  console.log(params);
  return request('/api/article/getone?id='+params.ID, {
    //params,
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
    },
    credentials: 'include',
  });
}

export async function getCollection(params) {
  return request('/api/articlefavorite/getfavoriteofarticle?articleID='+params.ID, {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
    },
    credentials: 'include',
  });
}

export async function getReply(params) {
  return request('/api/articlecomment/getbyarticleid?id='+params.ID+"&pageno="+params.PageNO+"&pagesize="+params.PageSize, {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
    },
    credentials: 'include',
  });
}

export async function getLike(params) {
  return request('/api/articlelike/getlikelist?articleID='+params.ID, {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
    },
    credentials: 'include',
  });
}

