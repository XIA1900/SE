/* Group page
Display contents:
1. Group name
2. Group introduction
3. list of posts as link, latest/hottest
4. button of join/delete group
5. Group avatar
6. number of group members
*/

import { LikeOutlined, LoadingOutlined, MessageOutlined, StarOutlined } from '@ant-design/icons';
import { Button, Card, Col, Form, List, Row, Select, Tag } from 'antd';
import React from 'react';
import { useRequest, history } from 'umi';
import ArticleListContent from './components';
import StandardFormRow from '@/pages/homepage/components/StandardFormRow';
import { getGroup } from '@/services/getGroupInfo';
import styles from './style.less';

const { Option } = Select;
const FormItem = Form.Item;
const pageSize = 10;
const groupName = history.location.search.substring(1);
console.log("groupName");
console.log(groupName);

const Posts = () => {
  console.log("going");
  const [form] = Form.useForm();
  
  const { data, reload, loading, loadMore, loadingMore } = useRequest(
    () => {
      console.log("going2");
      return getGroup({
        count: pageSize,
        type: 'hottest',
        groupName: groupName,
      });
    },
    {
      loadMore: true,
    },
  );
  const lists = data?.list[0] || [];
  console.log("lists");
  console.log(lists);
  const description = lists.groupDesciption;
  console.log(description);
  const avatar = lists.groupAvatar;
  const time = lists.createdAt;
  const members = lists.groupMember;
  const list = lists?.postList || [];
  // const list = lists.postList;
  // console.log(list);

  const IconText = ({ type, text }) => {
    switch (type) {
      case 'star-o':
        return (
          <span>
            <StarOutlined
              style={{
                marginRight: 8,
              }}
            />
            {text}
          </span>
        );

      case 'like-o':
        return (
          <span>
            <LikeOutlined
              style={{
                marginRight: 8,
              }}
            />
            {text}
          </span>
        );

      case 'message':
        return (
          <span>
            <MessageOutlined
              style={{
                marginRight: 8,
              }}
            />
            {text}
          </span>
        );

      default:
        return null;
    }
  };

  const formItemLayout = {
    wrapperCol: {
      xs: {
        span: 24,
      },
      sm: {
        span: 24,
      },
      md: {
        span: 12,
      },
    },
  };

  const loadMoreDom = list.length > 0 && (
    <div
      style={{
        textAlign: 'center',
        marginTop: 16,
      }}
    >
      <Button
        onClick={loadMore}
        style={{
          paddingLeft: 48,
          paddingRight: 48,
        }}
      >
        {loadingMore ? (
          <span>
            <LoadingOutlined /> Loading...
          </span>
        ) : (
          'Load More'
        )}
      </Button>
    </div>
  );

  return (
    <>
      <Card bordered={false}>
        <Form
          layout="inline"
          form={form}
          initialValues={{
            //owner: ['wjh', 'zxx'],
          }}
          onValuesChange={reload}
        >
          <StandardFormRow block> 
            <h1>{groupName}</h1>
            <p> {members} members</p>
            <p> Created at {time}</p>
          </StandardFormRow> 
          <p>{description}</p>

        </Form>
      </Card>
      <Card
        style={{
          marginTop: 24,
        }}
        bordered={false}
        bodyStyle={{
          padding: '8px 32px 32px 32px',
        }}
      >
        <List
          size="large"
          loading={loading}
          rowKey="id"
          itemLayout="vertical"
          loadMore={loadMoreDom}
          dataSource={list}
          renderItem={(item) => (
            <List.Item
              key={item.id}
              actions={[
                <IconText key="collection" type="star-o" text={item.collection} />,
                <IconText key="like" type="like-o" text={item.like} />,
                <IconText key="reply" type="message" text={item.reply} />,
              ]}
              extra={<div className={styles.listItemExtra} />}
            >
              <List.Item.Meta
                title={
                  <a className={styles.listItemMetaTitle} href={item.href}>
                    {item.title}
                  </a>
                }
              />
              <ArticleListContent data={item} />
            </List.Item>
          )}
        />
      </Card>
    </>
  );
};

export default Posts;
