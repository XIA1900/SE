import { Avatar } from 'antd';
import React from 'react';
import moment from 'moment';
import styles from './index.less';

const ArticleListContent = ({ data: { content, createdAt, owner, avatar} }) => (
  <div className={styles.listContent}>
    <div className={styles.description}>{content}</div>
    <div className={styles.extra}>
      <img src={avatar} style={{ width: '25px', height: '25px', borderRadius: '25px' }} />
      <a href=''> {owner} </a>
      <em> {moment(createdAt).format('YYYY-MM-DD HH:mm')}</em>
    </div>
  </div>
);

export default ArticleListContent;
