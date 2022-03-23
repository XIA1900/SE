import { PlusOutlined } from '@ant-design/icons';
import { Button, Card, List, Typography } from 'antd';
import { PageContainer } from '@ant-design/pro-layout';
import { useRequest, history } from 'umi';
import { getCreatedGroup } from '@/services/getGroupInfo';
import styles from './style.less';
import React, { useCallback } from 'react';

const { Paragraph } = Typography;
// a user can create at most 5 groups
const userName = history.location.search.substring(1);


const CardList = () => {
  const { data, loading } = useRequest(() => {
    return getCreatedGroup({
      userName: userName,
    });
  });

  const list = data?.list || [];

  console.log(list);

  const content = (
    <div className={styles.pageHeaderContent}>
      <p>
        Please select a group.
      </p>
    </div>
  );

  const onMenuClick = useCallback(
    (event) => {
      const { key } = event;
      history.push('/form/createGroup');
    },
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
          dataSource={[nullData, ...list]}
          renderItem={(item) => {
            if (item && item.id) {
              return (
                <List.Item key={item.id}>
                  <Card
                    hoverable
                    className={styles.card}
                    actions={[<p key="option1"># Group Members: {item.numberOfMember}</p>, <p key="option2"># Posts: {item.numberOfPost}</p>,  <p key="option3">Created At: {item.createdAt}</p>]}
                  >
                    <Card.Meta
                      avatar={<img alt="" className={styles.cardAvatar} src={item.groupAvatar} />}
                      title={<a href={"/group/management?"+item.groupName}>{item.groupName}</a>}
                      description={
                        <Paragraph
                          className={styles.item}
                          ellipsis={{
                            rows: 3,
                          }}
                        >
                          {item.groupDescription}
                        </Paragraph>
                      }
                    />
                  </Card>
                </List.Item>
              );
            }

            return (
              <List.Item>
                <Button type="dashed" className={styles.newButton} key="create" onClick = {onMenuClick}>
                  <PlusOutlined /> Create
                </Button>
              </List.Item>
            );
          }}
        />
      </div>
    </PageContainer>
  );
};

export default CardList;
