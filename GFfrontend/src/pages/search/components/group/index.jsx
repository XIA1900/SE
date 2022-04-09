import {
  ContactsOutlined,
  LikeOutlined,
  LoadingOutlined,
  MessageOutlined,
  StarOutlined,
} from '@ant-design/icons';
import { Button, Card, Col, Form, List, Row, Select, Tag, Tabs } from 'antd';
import React from 'react';
import { useRequest, history } from 'umi';
import ArticleListContent from '@/pages/group/content/components/articleContent';
import StandardFormRow from '@/pages/homepage/components/StandardFormRow';
import { searchGroup } from '@/services/search';
import styles from './style.less';

const { Option } = Select;
const FormItem = Form.Item;
const pageSize = 10;
const pageNo = 1;
const search = history.location.search.substring(1);


const Group = () => {
  const { data, loading } = useRequest( async () => {
    const result = await getJoinedGroup({

    });
    return result;
  },
  {
    formatResult: result => result,
  }
  );

  console.log(data);
  let list =  [];
  if(typeof(data) != 'undefined') {
    list = data;
  }

  const content = (
    <div className={styles.pageHeaderContent}>
      <p>
        Please select a group.
      </p>
    </div>
  );

  const nullData = {};
  return (
    <PageContainer content={content} >
      <div className={styles.cardList}>
        <List
          rowKey="id"
          loading={loading}
          grid={{
            gutter: 16,
            xs: 1,
            sm: 2,
            md: 3,
            lg: 3,
            xl: 4,
            xxl: 4,
          }}
          dataSource={[ ...list]}
          renderItem={(item) => {
            if (item && item.ID) {
              return (
                <List.Item key={item.ID}>
                  <Card
                    hoverable
                    className={styles.card}
                  >
                    <Card.Meta
                      avatar={<img alt="" className={styles.cardAvatar} src={item.groupAvatar} />}
                      title={<p key='group' onClick={() => {
                        history.push({
                          pathname: '/group/content',
                          search: item.ID,
                        });
                      }}>{item.Name}</p>}
                      description={
                        <Paragraph
                          className={styles.item}
                          ellipsis={{
                            rows: 3,
                          }}
                        >
                          {item.Description}
                          //<p>Created At: {item.CreateDay}</p>
                        </Paragraph>
                      }
                    />
                  </Card>
                </List.Item>
              );
            }

            return (
              <List.Item>
              </List.Item>
            );
          }}
        />
      </div>
    </PageContainer>
  );
};

export default Group;