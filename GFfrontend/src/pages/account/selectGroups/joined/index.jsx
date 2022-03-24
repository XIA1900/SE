import { PlusOutlined } from '@ant-design/icons';
import { Button, Card, List, Typography } from 'antd';
import { PageContainer } from '@ant-design/pro-layout';
import { useRequest, history } from 'umi';
import { getJoinedGroup } from '@/services/getGroupInfo';
import styles from './style.less';

const { Paragraph } = Typography;
const user = history.location.search;
const userName = user.substring(1);


const CardList = () => {
  const { data, loading } = useRequest(() => {
    return getJoinedGroup({
      userName: userName,
    });
  });

  const list = data?.list || [];
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
            if (item && item.id) {
              return (
                <List.Item key={item.id}>
                  <Card
                    hoverable
                    className={styles.card}
                    actions={[<p># Group Members: {item.numberOfMember}</p>, <p># Posts: {item.numberOfPost}</p>,  <p>Created At: {item.createdAt}</p>]}
                  >
                    <Card.Meta
                      avatar={<img alt="" className={styles.cardAvatar} src={item.groupAvatar} />}
                      title={<p key='group' onClick={() => {
                        history.push({
                          pathname: '/group/content',
                          search: item.groupName,
                        });
                      }}>{item.groupName}</p>}
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
              </List.Item>
            );
          }}
        />
      </div>
    </PageContainer>
  );
};

export default CardList;
