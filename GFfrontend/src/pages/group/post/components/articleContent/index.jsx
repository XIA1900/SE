import { Avatar } from 'antd';
import React from 'react';
import moment from 'moment';
import styles from './index.less';

const ArticleListContent = ({ data: { Content, CreateDay, Username, avatar} }) => (
  <div className={styles.listContent}>
    <div className={styles.extra}>
      <img src={avatar} style={{ width: '25px', height: '25px', borderRadius: '25px' }} />
      <a href=''> {Username} </a>
      <em> {moment(CreateDay).format('YYYY-MM-DD HH:mm')}</em>
    </div>
    <div className={styles.description}>{Content}</div>
  </div>
);

export default ArticleListContent;