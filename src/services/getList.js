import { request } from 'umi';

/* 
query list from homepage;

params = {count: pageSize} : pageSize(how many posts are displayed. eg: 10) 

return: a list, including pageSize posts and each post should have:
1. postid
2. owner name and href of personal center
3. title of the post, href of the post
4. first 30 words of post content
5. create date of the post
6. posted group and group link
7. the number of collections, number of likes, number of replies
*/

export async function queryList(params) {
    return request('/api/queryList', {
      params,
    });
  }
  