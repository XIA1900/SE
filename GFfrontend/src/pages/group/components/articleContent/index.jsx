import { Avatar } from 'antd';
import React from 'react';
import moment from 'moment';
import styles from './index.less';

const ArticleListContent = ({ data: { content, logo, updatedAt, owner, owner_href } }) => (
  <div className={styles.listContent}>
    <div className={styles.description}>{content}</div>
    <div className={styles.extra}>
      <img src={logo} style={{ width: '25px', height: '25px', borderRadius: '25px' }} />
      <a href={owner_href}> {owner} </a>
      <em>last updated at {moment(updatedAt).format('YYYY-MM-DD HH:mm')}</em>
    </div>
  </div>
);

export default ArticleListContent;
