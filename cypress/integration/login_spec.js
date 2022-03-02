

it('greets with Gator Forum', () => {
    cy.visit('http://localhost:8000/user/login/index.jsx')  // specify full URL if baseUrl is null or the domain is different the baseUrl
    cy.contains('ant-pro-form-login-title', 'Gator Forum');
})

it('explains with Gator Forum is a project created by Road Center', () => {
    cy.contains('ant-pro-form-login-desc', 'Gator Forum is a project created by Road Center');
})

it('links to https://github.com/fongziyjun16/SE', () => {
    cy
    .contains('Gator Forum')
    .should('have.attr','href','https://github.com/fongziyjun16/SE')
})

it('links to https://github.com/fongziyjun16/SE', () => {
    cy
    .contains('Road Center')
    .should('have.attr','href','https://github.com/fongziyjun16/SE')
})


it('requires UFID', () => {
    cy.get('form').contains('Login').click()
    cy.get('.error-messages')
      .should('contain','ufid can\'t be blank')
})

it('requires password', () => {
    cy.get('form').contains('Login').click()
    cy.get('.error-messages')
      .should('contain','ufid can\'t be blank')
})

it('navigates to #/homepage on successful login', () => {
    cy.get('[data-test=ufid]').type('28267331')
    cy.get('[data-test=password]').type('dfaewrq')
    cy.hash().should('eq', '#/homepage')
})