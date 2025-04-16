ALTER TABLE resolved_questions
    ALTER question_answer SET NOT NULL;
---- create above / drop below ----
ALTER TABLE resolved_questions
    ALTER question_answer DROP NOT NULL;
