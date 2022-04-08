import { Avatar } from 'antd';
import React from 'react';
import moment from 'moment';
import styles from './index.less';

const ArticleListContent = ({ data: { Content, UpdatedAt, Owner } }) => (
  <div className={styles.listContent}>
    <div className={styles.description}>{Content}</div>
    <div className={styles.extra}>
      <img src={'http://10.20.0.169:10010/resources/userfiles/'+Owner+'/avatar.png'} style={{ width: '25px', height: '25px', borderRadius: '25px' }} />
      <a href={'/account/center?'+Owner}> {Owner} </a>
      <em>last updated at {moment(UpdatedAt).format('YYYY-MM-DD HH:mm')}</em>
    </div>
  </div>
);

export default ArticleListContent;
