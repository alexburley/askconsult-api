import { mock } from 'jest-mock-extended'
import { TestAppCtx } from '../../../../test/test-manager'
import { SignupUserCommand } from './command'
import { UserRepository } from '../../../../adapters/repositories/user'
import { Email } from '../../../models/email'
import { UserDummy } from '../../../entities/test/dummy'

const repository = mock<UserRepository>()
const command = new SignupUserCommand(TestAppCtx(), {
  users: repository,
})

test('should signup a user', async () => {
  const result = await command.execute({
    name: 'John Doe',
    email: new Email('john@mail.com'),
    password: 'foobar',
  })

  expect(repository.persist).toHaveBeenCalledWith(result)
  expect(result).toEqual(UserDummy({ id: expect.stringContaining('user_') }))
})
